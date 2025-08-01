package api

import (
	"asyncflow/flowsvr/internal/err"
	"asyncflow/flowsvr/internal/service"
	"asyncflow/pkg/constant"
	"asyncflow/pkg/dto"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func CreateTask(c *gin.Context) {
	var req dto.CreateTaskReq
	// 1. 解析请求
	if e := c.ShouldBindJSON(&req); e != nil {
		log.Errorf("bind create task json failed: %+v", e)
		httpResponse := dto.BuildErrorResponse(constant.RspCodeInvalidInputErr)
		c.JSON(http.StatusOK, httpResponse)
		return
	}
	// 2. 执行业务逻辑
	rsp, e := service.CreateTask(&req)
	if e != nil {
		log.Errorf("handle create task failed: %+v", e)
		repCode := err.MapToRspCode(e.Code)
		httpResponse := dto.BuildErrorResponse(repCode)
		c.JSON(http.StatusOK, httpResponse)
		return
	}
	// 3. 返回结果
	if rsp != nil {
		httpResponse := dto.BuildSuccessResponse(rsp)
		c.JSON(http.StatusOK, httpResponse)
	} else {
		log.Error("create task response is nil")
		httpResponse := dto.BuildErrorResponse(constant.RspCodeInternalErr)
		c.JSON(http.StatusOK, httpResponse)
	}
}
