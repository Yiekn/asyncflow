package model

import (
	"time"
)

type TaskCfg struct {
	TaskType         string    `gorm:"column:task_type;type:varchar(128);primaryKey;comment:任务类型"`      // 任务类型
	TaskStage        int32     `gorm:"column:task_stage;type:tinyint unsigned;primaryKey;comment:任务阶段"` // 任务阶段
	Priority         int32     `gorm:"column:priority;type:int;not null;comment:优先级，单位为秒"`              // 优先级，单位为秒
	MaxRetryNum      int32     `gorm:"column:max_retry_num;type:int;not null;comment:最大重试次数"`           // 最大重试次数
	RetryInterval    int32     `gorm:"column:retry_interval;type:int;not null;comment:重试间隔，单位为秒"`       // 重试间隔，单位为秒
	MaxRetryInterval int32     `gorm:"column:max_retry_interval;type:int;not null;comment:最大重试间隔，单位为秒"` // 最大重试间隔，单位为秒
	CreateTime       time.Time `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	ModifyTime       time.Time `gorm:"column:modify_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
}

// TableName TaskCfg's table name
func (*TaskCfg) TableName() string {
	return "task_cfg"
}
