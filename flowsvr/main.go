package main

import (
	"asyncflow/flowsvr/initialize"
	"asyncflow/flowsvr/util"
	"github.com/gin-gonic/gin"
)

var logger = util.Logger

func main() {
	r := gin.Default()
	initialize.RegisterRouter(r)
	initialize.InitResource()
	err := r.Run()
	if err != nil {
		logger.Fatalf("run server failed, err: %s", err)
	}
}
