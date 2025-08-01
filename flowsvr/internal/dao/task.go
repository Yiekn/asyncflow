package dao

import (
	"asyncflow/flowsvr/internal/err"
	"asyncflow/flowsvr/internal/model"
	"asyncflow/flowsvr/internal/utils"
	"fmt"
	"gorm.io/gorm"
)

func CreateTask(tx *gorm.DB, task model.Task) *err.BizError {
	tableName := utils.ConvertToTableName(task.TaskId)
	if e := tx.Table(tableName).Create(&task).Error; e != nil {
		return err.NewBizErrorWithCause(err.BizCodeCreateTaskErr, fmt.Sprintf("create task failed, task_id=%s", task.TaskId), e)
	}
	return nil
}
