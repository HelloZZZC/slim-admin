package router

import (
	"github.com/gin-gonic/gin"
	"slim-admin/api"
)

type HealthRouter struct {
	Engine *gin.Engine
}

func (hr HealthRouter) register() {
	hr.Engine.GET("/health", new(api.HealthController).Health)
}
