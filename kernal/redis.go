package kernal

import (
	"github.com/redis/go-redis/v9"
	"slim-admin/global"
)

// initializeRDB 初始化redis客户端
func initializeRDB(ch chan *redis.Client) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     global.ApplicationConfig.Redis.Addr,
		Password: global.ApplicationConfig.Redis.Password,
		DB:       global.ApplicationConfig.Redis.DB,
	})

	ch <- rdb
}
