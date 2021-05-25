package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/winnix/go-apaleo-one-boilerplate/api/apaproxy"
	"github.com/winnix/go-apaleo-one-boilerplate/api/web"
)

type homeRequest struct {
	AccountCode string `form:"accountCode" validate:"required"`
	SubjectId   string `form:"subjectId" validate:"required"`
}

func Home(ctx *web.Context) web.HandlerResult {
	model := homeRequest{}

	if err := ctx.GinCtx.BindQuery(&model); err != nil {
		ctx.Logger.Sugar().Infof("failed binding request model: %s\n", err.Error())
		return BadRequest("invalid model")
	}

	if err := ctx.Validator.Struct(model); err != nil {
		ctx.Logger.Sugar().Infof("invalid request model: %s\n", err.Error())
		return BadRequest("invalid model")
	}

	proxy, err := apaproxy.NewApaProxy(ctx)
	if err != nil {
		ctx.Logger.Sugar().Errorf("failed initializing apaleo proxy: %s\n", err.Error())
		return ServerError()
	}

	reservations, err := proxy.BookingProxy.GetReservations()
	if err != nil {
		ctx.Logger.Sugar().Errorf("failed getting reservations: %s\n", err.Error())
		return ServerError()
	}

	return Render("home", gin.H{"Reservations": reservations})
}
