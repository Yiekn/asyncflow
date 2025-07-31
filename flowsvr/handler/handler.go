package handler

import (
	"asyncflow/flowsvr/db"
	"asyncflow/flowsvr/util"
	"gorm.io/gorm"
)

type Handler interface {
	HandleInput(tx *gorm.DB) error
	HandleProcess(tx *gorm.DB) error
}

func Run(hd Handler) {
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := hd.HandleInput(tx); err != nil {
		logger.Error(err)
		tx.Rollback()
		return
	}
	if err := hd.HandleProcess(tx); err != nil {
		logger.Error(err)
		tx.Rollback()
		return
	}
	tx.Commit()
}

var logger = util.Logger
