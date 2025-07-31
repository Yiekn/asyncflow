package db

import (
	"asyncflow/flowsvr/db/model"
	"fmt"
	"gorm.io/gorm"
)

func GetSchedulePos(tx *gorm.DB, taskType string) (model.SchedulePos, error) {
	var schedulePos model.SchedulePos
	if err := tx.Where("task_type = ?", taskType).First(&schedulePos).Error; err != nil {
		return model.SchedulePos{}, fmt.Errorf("schedule_pos not found, task_type=%s", taskType)
	}
	return schedulePos, nil
}
