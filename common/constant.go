package common

type TaskStatus = int32

const (
	TaskStatusWaiting   TaskStatus = 0
	TaskStatusExecuting TaskStatus = 1
	TaskStatusSucceed   TaskStatus = 2
	TaskStatusFailed    TaskStatus = 3
)

type RespCode = int32

const (
	Success                RespCode = 8000
	ErrInvalidInput        RespCode = 9001
	ErrTaskCfgNotFound     RespCode = 9002
	ErrSchedulePosNotFound RespCode = 9003
	ErrCreateTask          RespCode = 9004
	ErrUnknown             RespCode = 9999
)

var RespMsg = map[RespCode]string{
	Success:                "success",
	ErrInvalidInput:        "invalid input",
	ErrTaskCfgNotFound:     "task_cfg not found",
	ErrSchedulePosNotFound: "schedule_pos not found",
	ErrCreateTask:          "create task failed",
	ErrUnknown:             "unknown error",
}
