package kernal

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"slim-admin/global"
	"time"
)

// initializeConn 初始化数据库连接
func initializeConn() {
	global.GormDB = conn()
}

func conn() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@%s", global.ApplicationConfig.Mysql.Username, global.ApplicationConfig.Mysql.Password,
		global.ApplicationConfig.Mysql.Url)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	if global.ApplicationConfig.Mysql.ConnectionPool.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(global.ApplicationConfig.Mysql.ConnectionPool.MaxIdleConns)
	}
	if global.ApplicationConfig.Mysql.ConnectionPool.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(global.ApplicationConfig.Mysql.ConnectionPool.MaxOpenConns)
	}
	if global.ApplicationConfig.Mysql.ConnectionPool.ConnMaxLifeTimeSeconds > 0 {
		sqlDB.SetConnMaxLifetime(time.Second * time.Duration(global.ApplicationConfig.Mysql.ConnectionPool.
			ConnMaxLifeTimeSeconds))
	}

	return db
}
