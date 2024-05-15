package kernal

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"slim-admin/global"
)

// Bootstrap 服务初始化启动
func Bootstrap() {
	loadConfigYaml()

	cMysql := make(chan *gorm.DB)
	cRedis := make(chan *redis.Client)
	go initializeConn(cMysql)
	go initializeRDB(cRedis)
	global.GormDB = <-cMysql
	global.RDB = <-cRedis

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
