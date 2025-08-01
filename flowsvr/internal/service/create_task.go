package service

import (
	"asyncflow/flowsvr/internal/dao"
	"asyncflow/flowsvr/internal/err"
	"asyncflow/flowsvr/internal/model"
	"asyncflow/flowsvr/internal/utils"
	"asyncflow/pkg/constant"
	"asyncflow/pkg/dto"
	"gorm.io/gorm"
)

func CreateTask(req *dto.CreateTaskReq) (*dto.CreateTaskRsp, *err.BizError) {
	tx := utils.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	rsp, e := handle(req, tx)
	if e != nil {
		tx.Rollback()
		return nil, e
	}
	tx.Commit()
	return rsp, nil
}

func handle(req *dto.CreateTaskReq, tx *gorm.DB) (*dto.CreateTaskRsp, *err.BizError) {
	// 1. 校验必要参数
	if req.UserId == "" || req.TaskType == "" || req.TaskStage == nil {
		return nil, err.NewBizError(err.BizCodeInvalidInputErr, "lack of necessary params")
	}
	// 2. 获取任务配置
	taskCfg, e := dao.GetTaskCfg(tx, req.TaskType, *req.TaskStage)
	if e != nil {
		return nil, err.WrapBizError(e, "task configuration not found")
	}
	// 3. 填充任务信息
	task := model.Task{
		UserId:      req.UserId,
		TaskType:    req.TaskType,
		TaskStage:   *req.TaskStage,
		Status:      constant.TaskStatusWaiting,
		CrtRetryNum: 0,
		ScheduleLog: req.ScheduleLog,
		TaskContext: req.TaskContext,
	}
	// 3.1 优先使用用户指定的参数
	if req.Priority != nil {
		task.Priority = *req.Priority
	} else {
		task.Priority = taskCfg.Priority
	}
	if req.MaxRetryNum != nil {
		task.MaxRetryNum = *req.MaxRetryNum
	} else {
		task.MaxRetryNum = taskCfg.MaxRetryNum
	}
	if req.RetryInterval != nil {
		task.RetryInterval = *req.RetryInterval
	} else {
		task.RetryInterval = taskCfg.RetryInterval
	}
	if req.MaxRetryInterval != nil {
		task.MaxRetryInterval = *req.MaxRetryInterval
	} else {
		task.MaxRetryInterval = taskCfg.MaxRetryInterval
	}
	// 3.2 计算并设置调度时间
	task.OrderTime = task.CalculateOrderTime(false)
	// 3.3 生成任务ID
	// 3.3.1 获取调度位置
	schedulePos, e := dao.GetSchedulePos(tx, task.TaskType)
	if e != nil {
		return nil, err.WrapBizError(e, "schedule position not found")
	}
	// 3.3.2 生成任务ID
	snowflakeId, ce := utils.GetSnowflakeGenerator().NextId()
	if ce != nil {
		return nil, err.NewBizErrorWithCause(err.BizCodeSnowflakeErr, "snowflake generate failed", e)
	}
	// 3.3.3 设置任务ID
	task.TaskId = utils.JointTaskId(snowflakeId, schedulePos.GetEndTableName())
	// 4. 插入任务
	e = dao.CreateTask(tx, task)
	if e != nil {
		return nil, err.WrapBizError(e, "failed to create task")
	}
	// 5. 设置响应数据
	rsp := &dto.CreateTaskRsp{
		TaskId: task.TaskId,
	}
	return rsp, nil
}
