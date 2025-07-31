package common

type BaseResp struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func GetRespBase(code RespCode) BaseResp {
	return BaseResp{
		Code:    code,
		Message: RespMsg[code],
	}
}

type CreateTaskReq struct {
	UserId           string `json:"userId"`
	TaskType         string `json:"taskType"`
	TaskStage        *int32 `json:"taskStage"`
	Priority         *int32 `json:"priority"`
	MaxRetryNum      *int32 `json:"maxRetryNum"`
	RetryInterval    *int32 `json:"retryInterval"`
	MaxRetryInterval *int32 `json:"maxRetryInterval"`
	ScheduleLog      string `json:"scheduleLog"`
	TaskContext      string `json:"taskContext"`
}

type CreateTaskResp struct {
	BaseResp
	TaskId string `json:"taskId"`
}
