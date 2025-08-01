package constant

type TaskStatus = int32

const (
	TaskStatusWaiting   TaskStatus = 0
	TaskStatusExecuting TaskStatus = 1
	TaskStatusSucceed   TaskStatus = 2
	TaskStatusFailed    TaskStatus = 3
)

type RspCode int32

const (
	RspCodeSuccess                RspCode = 8000
	RspCodeInvalidInputErr        RspCode = 9001
	RspCodeTaskCfgNotFoundErr     RspCode = 9002
	RspCodeSchedulePosNotFoundErr RspCode = 9003
	RspCodeCreateTaskErr          RspCode = 9004
	RspCodeSnowflakeErr           RspCode = 9101
	RspCodeInternalErr            RspCode = 9999
)

var RspMsgDict = map[RspCode]string{
	RspCodeSuccess:                "success",
	RspCodeInvalidInputErr:        "invalid input",
	RspCodeTaskCfgNotFoundErr:     "task_cfg not found",
	RspCodeSchedulePosNotFoundErr: "schedule_pos not found",
	RspCodeCreateTaskErr:          "create task failed",
	RspCodeSnowflakeErr:           "snowflake generate failed",
	RspCodeInternalErr:            "internal error",
}
