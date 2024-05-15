package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"slim-admin/api"
	"slim-admin/config"
)

var (
	GormDB            *gorm.DB
	Controllers       api.Controllers
	Viper             *viper.Viper
	ApplicationConfig config.Application
	RDB               *redis.Client
)
