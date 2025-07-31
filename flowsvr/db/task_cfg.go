package db

import (
	"asyncflow/flowsvr/db/model"
	"fmt"
	"gorm.io/gorm"
)

func GetTaskCfg(tx *gorm.DB, taskType string, taskStage int32) (model.TaskCfg, error) {
	var taskCfg model.TaskCfg
	if err := tx.Where("task_type = ? AND task_stage = ?", taskType, taskStage).First(&taskCfg).Error; err != nil {
		return model.TaskCfg{}, fmt.Errorf("task_cfg not found, task_type=%s, task_stage=%d", taskType, taskStage)
	}
	return taskCfg, nil
}
