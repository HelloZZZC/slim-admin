package kernal

import (
	"github.com/gin-gonic/gin"
	"slim-admin/router"
)

// registerRoute 注册路由
func registerRoute(e *gin.Engine) *gin.Engine {
	e.Group("")
	{
		router.HealthRouter{E: e}.Register()
		router.StaffRouter{E: e}.Register()
		router.ExampleRouter{E: e}.Register()
	}

	return e
}
