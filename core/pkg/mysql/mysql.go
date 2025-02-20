package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	mysqlConfig "sweete/core/pkg/mysql/config"
	"sweete/user_service/config"
	"sync"
	"time"
)

const (
	maxConn           = 50
	maxIdConn         = 10
	maxConnLifetime   = 3 * time.Minute
	maxIdConnLifetime = 1 * time.Minute
	lazyConnect       = false
)

var instance *gorm.DB
var once sync.Once

func GetInstance() *gorm.DB {
	once.Do(func() {
		if instance == nil {
			universalClient, err := NewGormDB(config.GetInstance().MySql)
			if err != nil {
				panic(err)
			}
			instance = universalClient
		}
	})
	return instance
}

func NewGormDB(config *mysqlConfig.Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.Url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	return db, err
}
