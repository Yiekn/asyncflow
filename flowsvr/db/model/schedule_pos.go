package model

import (
	"fmt"
	"time"
)

type SchedulePos struct {
	TaskType         string    `gorm:"column:task_type;type:varchar(128);primaryKey;comment:任务类型"`   // 任务类型
	ScheduleBeginPos int32     `gorm:"column:schedule_begin_pos;type:int;not null;comment:调度开始于几号表"` // 调度开始于几号表
	ScheduleEndPos   int32     `gorm:"column:schedule_end_pos;type:int;not null;comment:调度结束于几号表"`   // 调度结束于几号表
	CreateTime       time.Time `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	ModifyTime       time.Time `gorm:"column:modify_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
}

// TableName SchedulePo's table name
func (*SchedulePos) TableName() string {
	return "schedule_pos"
}

func (pos *SchedulePos) GetBeginTableName() string {
	return fmt.Sprintf("task_%s_%d", pos.TaskType, pos.ScheduleBeginPos)
}

func (pos *SchedulePos) GetEndTableName() string {
	return fmt.Sprintf("task_%s_%d", pos.TaskType, pos.ScheduleEndPos)
}
