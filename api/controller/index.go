package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/winnix/go-apaleo-one-boilerplate/api/web"
)

func Index(ctx *web.Context) web.HandlerResult {
	params := gin.H{
		"installUrl": fmt.Sprintf("%s/install", ctx.Config.Host),
	}
	return Render("index", params)
}
