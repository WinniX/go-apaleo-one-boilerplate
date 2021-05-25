package web

import (
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/openidConnect"
	"github.com/winnix/go-apaleo-one-boilerplate/api/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	SessionName           = "myapp"
	CurrentUserSessionKey = "currentuser"
	RedirectURLSessionKey = "redirecturl"
)

var (
	sessionOptions     = sessions.Options{Path: "/", SameSite: http.SameSiteNoneMode, Secure: true}
	authSessionOptions = sessions.Options{SameSite: http.SameSiteNoneMode, Secure: true}
)

type Context struct {
	Db        *gorm.DB
	Config    *config.Config
	Ctx       context.Context
	GinCtx    *gin.Context
	Validator *validator.Validate
	Logger    *zap.Logger
	User      *CurrentUser
}

type CurrentUser struct {
	ID          uuid.UUID
	AccountCode string
}

func (ctx *Context) SetCurrentUser(id uuid.UUID, accountCode string) {
	ctx.User = &CurrentUser{
		ID:          id,
		AccountCode: accountCode,
	}

	StoreInSession(ctx.GinCtx, CurrentUserSessionKey, ctx.User)
}

func (ctx *Context) SetCurrentUserIfExists() {
	user, err := GetCurrentUser(ctx.GinCtx)
	if err != nil {
		return
	}

	ctx.SetCurrentUser(user.ID, user.AccountCode)
}

func GetCurrentUser(ginCtx *gin.Context) (CurrentUser, error) {
	session := sessions.Default(ginCtx)
	maybeUser := session.Get(CurrentUserSessionKey)
	if maybeUser == nil {
		return CurrentUser{}, errors.New("current user is not set")
	}

	user, ok := maybeUser.(CurrentUser)
	if !ok {
		return CurrentUser{}, errors.New("can't cast current user")
	}

	return user, nil
}

func InitSessions(router *gin.Engine, cfg *config.Config) {
	gob.Register(CurrentUser{})
	store := cookie.NewStore([]byte(cfg.SessionSecretKey))
	store.Options(sessionOptions)
	router.Use(sessions.Sessions(SessionName, store))
}

func InitAuthProviders(cfg *config.Config) error {
	store := cookie.NewStore([]byte(cfg.SessionSecretKey))
	store.Options(authSessionOptions)
	gothic.Store = store
	gothic.GetProviderName = func(req *http.Request) (string, error) {
		return "openid-connect", nil
	}

	apaOidc, err := GetApaleoAuthProvider(cfg)
	if err != nil {
		return err
	}

	goth.UseProviders(apaOidc)

	return nil
}

func GetApaleoAuthProvider(cfg *config.Config) (*openidConnect.Provider, error) {
	apaCfg := cfg.OAuth["apaleo"]
	apaOidc, err := openidConnect.New(
		apaCfg.ClientID,
		apaCfg.ClientSecret,
		fmt.Sprintf("%s/auth/apaleo/callback", cfg.Host),
		apaCfg.AutoDiscoveryUrl,
		apaCfg.Scopes...)
	if err != nil {
		return nil, err
	}

	return apaOidc, nil
}

func StoreInSession(ginCtx *gin.Context, key string, value interface{}) {
	session := sessions.Default(ginCtx)
	session.Set(key, value)
	session.Options(sessionOptions)
	session.Save()
}

func ExtractFromSession(ginCtx *gin.Context, key string) (interface{}, bool) {
	session := sessions.Default(ginCtx)
	if value := session.Get(key); value != nil {
		session.Delete(key)
		return value, true
	}

	return "", false
}

type HandlerResult interface {
	Render(c *gin.Context)
}

type JsonResult struct {
	Status int
	Json   interface{}
}

type HtmlResult struct {
	Status   int
	Template string
	Params   gin.H
}

func (r JsonResult) Render(c *gin.Context) {
	c.JSON(r.Status, gin.H{"result": r.Json})
}

func (r HtmlResult) Render(c *gin.Context) {
	c.HTML(r.Status, r.Template, r.Params)
}

type HandlerFunc func(*Context) HandlerResult
