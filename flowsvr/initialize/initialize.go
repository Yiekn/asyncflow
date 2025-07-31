package initialize

import (
	"asyncflow/flowsvr/config"
	"asyncflow/flowsvr/controller"
	"asyncflow/flowsvr/db"
	"asyncflow/flowsvr/util"
	"github.com/gin-gonic/gin"
)

var logger = util.Logger

func RegisterRouter(router *gin.Engine) {
	router.POST("/tasks/create", controller.CreateTask)
}

func InitResource() {
	config.InitConfig()
	err := db.InitDB()
	if err != nil {
		logger.Fatalf("%s", err)
	}
}
