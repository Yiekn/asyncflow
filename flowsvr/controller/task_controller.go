package controller

import (
	"asyncflow/common"
	"asyncflow/flowsvr/handler"
	"asyncflow/flowsvr/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

var logger = util.Logger

func CreateTask(c *gin.Context) {
	var hd handler.CreateTaskHandler
	defer func() {
		c.JSON(http.StatusOK, hd.Resp)
	}()
	// 1. 解析请求
	if err := c.ShouldBindJSON(&hd.Req); err != nil {
		logger.Errorf("should bind failed, err=%s", err)
		hd.Resp.BaseResp = common.GetRespBase(common.ErrInvalidInput)
		return
	}
	// 2. 执行逻辑
	handler.Run(&hd)
}
