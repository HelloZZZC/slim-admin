package router

import (
	"github.com/gin-gonic/gin"
	"slim-admin/global"
	"slim-admin/middleware"
)

type ExampleRouter struct {
	E *gin.Engine
}

func (er ExampleRouter) Register() {
	g := er.E.Group("/example").Use(middleware.PanicInterceptor())
	{
		g.GET("/bizError", global.Controllers.ExampleController.BizError)
	}
}
