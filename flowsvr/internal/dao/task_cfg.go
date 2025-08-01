package dao

import (
	"asyncflow/flowsvr/internal/err"
	"asyncflow/flowsvr/internal/model"
	"fmt"
	"gorm.io/gorm"
)

func GetTaskCfg(tx *gorm.DB, taskType string, taskStage int32) (model.TaskCfg, *err.BizError) {
	var taskCfg model.TaskCfg
	if e := tx.Where("task_type = ? AND task_stage = ?", taskType, taskStage).First(&taskCfg).Error; e != nil {
		msg := fmt.Sprintf("task_cfg not found, task_type=%s, task_stage=%d", taskType, taskStage)
		return model.TaskCfg{}, err.NewBizErrorWithCause(err.BizCodeTaskCfgNotFoundErr, msg, e)
	}
	return taskCfg, nil
}
