package router

import (
	"github.com/gin-gonic/gin"
	"slim-admin/global"
)

type HealthRouter struct {
	E *gin.Engine
}

func (hr HealthRouter) Register() {
	hr.E.GET("/health", global.Controllers.HealthController.Health)
}
