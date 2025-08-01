package utils

import (
	"asyncflow/flowsvr/config"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	db     *gorm.DB
	dbOnce sync.Once
)

func GetDB() *gorm.DB {
	dbOnce.Do(func() {
		initDB()
	})
	return db
}

func initDB() {
	conf := config.GetGlobalConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Mysql.Username,
		conf.Mysql.Password,
		conf.Mysql.Host,
		conf.Mysql.Port,
		conf.Mysql.Database)
	log.Debugf("dsn=%s", dsn)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("database connect failed, err: %s", err.Error())
	}
}
