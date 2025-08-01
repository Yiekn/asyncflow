package dto

import "asyncflow/pkg/constant"

type HttpResponse struct {
	Code    constant.RspCode `json:"code"`
	Message string           `json:"message"`
	Data    interface{}      `json:"data"`
}

func NewHttpResponse(code constant.RspCode, message string, data interface{}) *HttpResponse {
	return &HttpResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func BuildSuccessResponse(data interface{}) *HttpResponse {
	return &HttpResponse{
		Code:    constant.RspCodeSuccess,
		Message: constant.RspMsgDict[constant.RspCodeSuccess],
		Data:    data,
	}
}

func BuildErrorResponse(code constant.RspCode) *HttpResponse {
	return &HttpResponse{
		Code:    code,
		Message: constant.RspMsgDict[code],
		Data:    nil,
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

type CreateTaskRsp struct {
	TaskId string `json:"taskId"`
}
