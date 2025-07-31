package db

import (
	"asyncflow/flowsvr/db/model"
	"asyncflow/flowsvr/util"
	"fmt"
	"gorm.io/gorm"
)

func CreateTask(tx *gorm.DB, task model.Task) error {
	tableName := util.ConvertToTableName(task.TaskId)
	if err := tx.Table(tableName).Create(&task).Error; err != nil {
		return fmt.Errorf("create task failed, err=%s", err)
	}
	return nil
}
