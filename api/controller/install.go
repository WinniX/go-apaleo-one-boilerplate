package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/markbates/goth/gothic"
	"github.com/winnix/go-apaleo-one-boilerplate/api/apaproxy"
	"github.com/winnix/go-apaleo-one-boilerplate/api/domain"
	"github.com/winnix/go-apaleo-one-boilerplate/api/service"
	"github.com/winnix/go-apaleo-one-boilerplate/api/web"
	"gorm.io/gorm"
)

type apaleoClaims struct {
	UserId      uuid.UUID `json:"sub"`
	AccountCode string    `json:"account_code"`
}

func (c *apaleoClaims) Valid() error {
	if len(c.AccountCode) == 0 {
		return errors.New("account_code claim is empty")
	}

	if c.UserId == uuid.Nil {
		return errors.New("sub claim is empty")
	}

	return nil
}

func Install(ctx *web.Context) web.HandlerResult {
	ctx.GinCtx.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/auth/apaleo", ctx.Config.Host))
	return nil
}

func AuthApaleo(ctx *web.Context) web.HandlerResult {
	gothic.BeginAuthHandler(ctx.GinCtx.Writer, ctx.GinCtx.Request)
	return nil
}

func AuthApaleoCallback(ctx *web.Context) web.HandlerResult {
	user, err := gothic.CompleteUserAuth(ctx.GinCtx.Writer, ctx.GinCtx.Request)
	if err != nil {
		ctx.Logger.Sugar().Warnf("apaleo authorization failed: %s\n", err.Error())
		return ServerError()
	}

	claims := apaleoClaims{}
	jwt.ParseWithClaims(user.AccessToken, &claims, nil)
	if err = claims.Valid(); err != nil {
		ctx.Logger.Sugar().Errorf("error parsing access token: %s\n", err.Error())
		return ServerError()
	}

	userId := claims.UserId

	userTokenService := service.NewUserTokenService(ctx.Db, claims.AccountCode)
	userToken, err := userTokenService.GetCurrent(userId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.Logger.Sugar().Errorf("error getting current user token: %s\n", err.Error())
		return ServerError()
	}

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		userToken = &domain.UserToken{}
		userToken.UserUID = userId
		userToken.AccountCode = claims.AccountCode
	}

	userToken.AccessToken = user.AccessToken
	userToken.RefreshToken = user.RefreshToken
	userToken.Expiry = user.ExpiresAt
	if err := userTokenService.SaveUserToken(userToken); err != nil {
		ctx.Logger.Sugar().Errorf("failed saving user token: %s\n", err.Error())
		return ServerError()
	}

	ctx.SetCurrentUser(userId, userToken.AccountCode)

	proxy, err := apaproxy.NewApaProxy(ctx)
	if err != nil {
		ctx.Logger.Sugar().Errorf("failed initializing apaleo proxy: %s\n", err.Error())
		return ServerError()
	}

	id, err := proxy.IntegrationProxy.CreateIntegration(fmt.Sprintf("%s/home", ctx.Config.Host))
	if err != nil {
		ctx.Logger.Sugar().Errorf("failed saving ui integration: %s\n", err.Error())
		return ServerError()
	}

	redirectURL, exists := web.ExtractFromSession(ctx.GinCtx, web.RedirectURLSessionKey)
	if exists {
		if url, ok := redirectURL.(string); ok {
			ctx.GinCtx.Redirect(http.StatusTemporaryRedirect, url)
			ctx.GinCtx.AbortWithStatus(http.StatusTemporaryRedirect)
			return nil
		}

		ctx.Logger.Sugar().Errorf("redirect URL is invalid: %s\n", redirectURL)
	}

	params := gin.H{
		"appUrl": fmt.Sprintf("https://app.apaleo.com/apps/%s", id),
	}
	return Render("install", params)
}
