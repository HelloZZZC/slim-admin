package router

import (
	"github.com/gin-gonic/gin"
	"slim-admin/api"
)

type HealthRouter struct {
	engine *gin.Engine
}

func (hr HealthRouter) register() {
	hr.engine.GET("/health", api.Health)
}
