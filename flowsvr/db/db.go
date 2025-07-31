package db

import (
	"asyncflow/flowsvr/config"
	"asyncflow/flowsvr/util"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var logger = util.Logger

func InitDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.Mysql.Username,
		config.Config.Mysql.Password,
		config.Config.Mysql.Host,
		config.Config.Mysql.Port,
		config.Config.Mysql.Database)
	logger.Debugf("dsn=%s", dsn)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.New("database connect failed")
	}
	return nil
}
