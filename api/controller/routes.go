package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/winnix/go-apaleo-one-boilerplate/api/config"
	"github.com/winnix/go-apaleo-one-boilerplate/api/middleware"
	"github.com/winnix/go-apaleo-one-boilerplate/api/web"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Render(template string, params gin.H) web.HtmlResult {
	return web.HtmlResult{
		Status:   http.StatusOK,
		Template: fmt.Sprintf("%s.tmpl", template),
		Params:   params,
	}
}

func BadRequest(json string) web.JsonResult {
	return web.JsonResult{
		Status: http.StatusBadRequest,
		Json:   json,
	}
}

func ServerError() web.HtmlResult {
	return web.HtmlResult{
		Status:   http.StatusInternalServerError,
		Template: "500.tmpl",
	}
}

func getContext(c *gin.Context) *web.Context {
	ctx := &web.Context{
		Db:        c.MustGet("dsn").(*gorm.DB),
		Config:    c.MustGet("config").(*config.Config),
		Ctx:       context.Background(),
		GinCtx:    c,
		Validator: c.MustGet("validator").(*validator.Validate),
		Logger:    c.MustGet("logger").(*zap.Logger),
	}

	ctx.SetCurrentUserIfExists()

	return ctx
}

func wrapper(f web.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := getContext(c)

		res := f(ctx)
		if res == nil {
			return
		}

		res.Render(c)
	}
}

func SetRoutes(r *gin.Engine, cfg *config.Config) {
	r.GET("/", wrapper(Index))

	r.GET("/install", wrapper(Install))
	r.GET("/auth/apaleo", wrapper(AuthApaleo))
	r.GET("/auth/apaleo/callback", wrapper(AuthApaleoCallback))

	authorized := r.Group("/")
	authorized.Use(middleware.ApaleoAuthRequired(cfg))
	{
		authorized.GET("/home", wrapper(Home))
	}
}
