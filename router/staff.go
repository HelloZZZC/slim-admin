package router

import (
	"github.com/gin-gonic/gin"
	"slim-admin/global"
	"slim-admin/middleware"
)

type StaffRouter struct {
	E *gin.Engine
}

func (sr StaffRouter) Register() {
	g := sr.E.Group("/staff").Use(middleware.PanicInterceptor())
	{
		g.GET("/info", global.Controllers.StaffController.GetStaff)
	}
}
