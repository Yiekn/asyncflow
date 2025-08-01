package dao

import (
	"asyncflow/flowsvr/internal/err"
	"asyncflow/flowsvr/internal/model"
	"fmt"
	"gorm.io/gorm"
)

func GetSchedulePos(tx *gorm.DB, taskType string) (model.SchedulePos, *err.BizError) {
	var schedulePos model.SchedulePos
	if e := tx.Where("task_type = ?", taskType).First(&schedulePos).Error; e != nil {
		msg := fmt.Sprintf("schedule_pos not found, task_type=%s", taskType)
		return model.SchedulePos{}, err.NewBizErrorWithCause(err.BizCodeSchedulePosNotFoundErr, msg, e)
	}
	return schedulePos, nil
}
