package router

import (
	"asyncflow/flowsvr/api"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.POST("/tasks/create", api.CreateTask)
}
