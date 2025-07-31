package handler

import (
	"asyncflow/common"
	"asyncflow/flowsvr/db"
	"asyncflow/flowsvr/db/model"
	"asyncflow/flowsvr/util"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type CreateTaskHandler struct {
	Req  common.CreateTaskReq
	Resp common.CreateTaskResp
	task model.Task
}

func (h *CreateTaskHandler) HandleInput(tx *gorm.DB) error {
	// 1. 校验必要参数
	if h.Req.UserId == "" || h.Req.TaskType == "" || h.Req.TaskStage == nil {
		h.Resp.BaseResp = common.GetRespBase(common.ErrInvalidInput)
		return errors.New("lack of necessary params")
	}
	// 2. 获取任务配置
	taskCfg, err := db.GetTaskCfg(tx, h.Req.TaskType, *h.Req.TaskStage)
	if err != nil {
		h.Resp.BaseResp = common.GetRespBase(common.ErrTaskCfgNotFound)
		return err
	}
	// 3. 填充任务信息
	h.task = model.Task{
		UserId:      h.Req.UserId,
		TaskType:    h.Req.TaskType,
		TaskStage:   *h.Req.TaskStage,
		Status:      common.TaskStatusWaiting,
		CrtRetryNum: 0,
		ScheduleLog: h.Req.ScheduleLog,
		TaskContext: h.Req.TaskContext,
	}
	// 3.1 优先使用用户指定的参数
	if h.Req.Priority != nil {
		h.task.Priority = *h.Req.Priority
	} else {
		h.task.Priority = taskCfg.Priority
	}
	if h.Req.MaxRetryNum != nil {
		h.task.MaxRetryNum = *h.Req.MaxRetryNum
	} else {
		h.task.MaxRetryNum = taskCfg.MaxRetryNum
	}
	if h.Req.RetryInterval != nil {
		h.task.RetryInterval = *h.Req.RetryInterval
	} else {
		h.task.RetryInterval = taskCfg.RetryInterval
	}
	if h.Req.MaxRetryInterval != nil {
		h.task.MaxRetryInterval = *h.Req.MaxRetryInterval
	} else {
		h.task.MaxRetryInterval = taskCfg.MaxRetryInterval
	}
	// 3.2 计算并设置调度时间
	h.task.OrderTime = h.task.CalculateOrderTime(false)
	// 3.3 生成任务ID
	// 3.3.1 获取调度位置
	schedulePos, err := db.GetSchedulePos(tx, h.task.TaskType)
	if err != nil {
		h.Resp.BaseResp = common.GetRespBase(common.ErrSchedulePosNotFound)
		return err
	}
	// 3.3.2 生成任务ID
	snowflakeId, err := util.SnowflakeGenerator.NextId()
	if err != nil {
		h.Resp.BaseResp = common.GetRespBase(common.ErrUnknown)
		return fmt.Errorf("snowflake generate failed, err=%s", err)
	}
	h.task.TaskId = util.JointTaskId(snowflakeId, schedulePos.GetEndTableName())
	return nil
}

func (h *CreateTaskHandler) HandleProcess(tx *gorm.DB) error {
	// 1. 插入任务
	err := db.CreateTask(tx, h.task)
	if err != nil {
		h.Resp.BaseResp = common.GetRespBase(common.ErrCreateTask)
		return err
	}
	// 2. 设置响应
	h.Resp.BaseResp = common.GetRespBase(common.Success)
	h.Resp.TaskId = h.task.TaskId
	return nil
}
