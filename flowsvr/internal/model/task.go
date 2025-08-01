package model

import (
	"time"
)

type Task struct {
	Id               int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:主键"`    // 主键
	UserId           string    `gorm:"column:user_id;type:varchar(256);not null;comment:用户id，标识用户"`       // 用户id，标识用户
	TaskId           string    `gorm:"column:task_id;type:varchar(256);not null;comment:任务id，标识任务"`       // 任务id，标识任务
	TaskType         string    `gorm:"column:task_type;type:varchar(128);not null;comment:任务类型"`          // 任务类型
	TaskStage        int32     `gorm:"column:task_stage;type:tinyint unsigned;not null;comment:任务阶段"`     // 任务阶段
	Status           int32     `gorm:"column:status;type:tinyint unsigned;not null;comment:状态"`           // 状态
	Priority         int32     `gorm:"column:priority;type:int;not null;comment:优先级，单位为秒"`                // 优先级，单位为秒
	CrtRetryNum      int32     `gorm:"column:crt_retry_num;type:int;not null;comment:已经重试几次了"`            // 已经重试几次了
	MaxRetryNum      int32     `gorm:"column:max_retry_num;type:int;not null;comment:最大能重试几次"`            // 最大能重试几次
	RetryInterval    int32     `gorm:"column:retry_interval;type:int;not null;comment:重试间隔，单位为秒"`         // 重试间隔，单位为秒
	MaxRetryInterval int32     `gorm:"column:max_retry_interval;type:int;not null;comment:最大重试间隔，单位为秒"`   // 最大重试间隔，单位为秒
	ScheduleLog      string    `gorm:"column:schedule_log;type:varchar(4096);not null;comment:调度信息记录"`    // 调度信息记录
	TaskContext      string    `gorm:"column:task_context;type:varchar(8192);not null;comment:任务上下文"`     // 任务上下文
	OrderTime        int64     `gorm:"column:order_time;type:bigint;not null;comment:调度时间，越小调度越优先，单位为毫秒"` // 调度时间，越小调度越优先，单位为毫秒
	CreateTime       time.Time `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	ModifyTime       time.Time `gorm:"column:modify_time;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
}

// CalculateOrderTime 计算任务的调度时间
// 参数 withRetryTime 表示是否考虑重试时间
func (t *Task) CalculateOrderTime(withRetryTime bool) int64 {
	// 如果不考虑重试时间
	if !withRetryTime {
		// 返回当前时间减去优先级对应的毫秒数
		return time.Now().UnixMilli() - int64(t.Priority)*1000
	} else {
		// 返回当前时间加上重试时间
		return time.Now().UnixMilli() + t.calculateRetryTime()
	}
}

// calculateRetryTime 计算任务的重试时间
func (t *Task) calculateRetryTime() int64 {
	// 如果重试间隔为0，则返回0
	if t.RetryInterval == 0 {
		return 0
		// 如果重试间隔小于0，则返回重试间隔的绝对值
	} else if t.RetryInterval < 0 {
		return -int64(t.RetryInterval)
	} else {
		// 初始化当前重试时间为1
		curRetryTime := int64(1)
		// 根据当前重试次数计算重试时间
		for i := 0; i < int(t.CrtRetryNum); i++ {
			curRetryTime *= 2
		}
		// 计算最终的重试时间
		curRetryTime *= int64(t.RetryInterval)
		// 返回重试时间和最大重试间隔的最小值
		return min(curRetryTime, int64(t.MaxRetryInterval))
	}
}
