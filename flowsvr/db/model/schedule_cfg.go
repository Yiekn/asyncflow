package model

import (
	"time"
)

type ScheduleCfg struct {
	TaskType          string    `gorm:"column:task_type;type:varchar(128);primaryKey;comment:任务类型"`                 // 任务类型
	TaskStage         int32     `gorm:"column:task_stage;type:tinyint unsigned;primaryKey;comment:任务阶段"`            // 任务阶段
	ScheduleLimit     int32     `gorm:"column:schedule_limit;type:int;not null;comment:一次拉取多少个任务"`                  // 一次拉取多少个任务
	ScheduleInterval  int32     `gorm:"column:schedule_interval;type:int;not null;comment:拉取任务的间隔，单位为秒"`            // 拉取任务的间隔，单位为秒
	MaxProcessingTime int32     `gorm:"column:max_processing_time;type:int;not null;comment:Worker处于执行中的最大时间，单位为秒"` // Worker处于执行中的最大时间，单位为秒
	CreateTime        time.Time `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	ModifyTime        time.Time `gorm:"column:modify_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
}

// TableName ScheduleCfg's table name
func (*ScheduleCfg) TableName() string {
	return "schedule_cfg"
}
