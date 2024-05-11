package kernal

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func Conn() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@%s", ApplicationConfig.Mysql.Username, ApplicationConfig.Mysql.Password,
		ApplicationConfig.Mysql.Url)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	if ApplicationConfig.Mysql.ConnectionPool.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(ApplicationConfig.Mysql.ConnectionPool.MaxIdleConns)
	}
	if ApplicationConfig.Mysql.ConnectionPool.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(ApplicationConfig.Mysql.ConnectionPool.MaxOpenConns)
	}
	if ApplicationConfig.Mysql.ConnectionPool.ConnMaxLifeTimeSeconds > 0 {
		sqlDB.SetConnMaxLifetime(time.Second * time.Duration(ApplicationConfig.Mysql.ConnectionPool.
			ConnMaxLifeTimeSeconds))
	}

	return db
}
