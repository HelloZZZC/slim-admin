package router

import (
	"github.com/gin-gonic/gin"
	"slim-admin/global"
)

type StaffRouter struct {
	E *gin.Engine
}

func (sr StaffRouter) Register() {
	g := sr.E.Group("/staff")
	{
		g.GET("/info", global.Controllers.StaffController.GetStaff)
	}
}
