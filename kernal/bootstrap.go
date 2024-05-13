package kernal

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"slim-admin/global"
)

// Bootstrap 服务初始化启动
func Bootstrap() {
	loadConfigYaml()
	initializeConn()
	executeDI()
	engine := gin.Default()
	registerRoute(engine)

	var port int
	if port = global.ApplicationConfig.Server.Port; port == 0 {
		port = 8080
	}

	addr := fmt.Sprintf(":%d", port)
	err := endless.ListenAndServe(addr, engine)
	if err != nil {
		panic(err)
	}
}
