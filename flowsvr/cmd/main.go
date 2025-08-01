package main

import (
	"asyncflow/flowsvr/config"
	"asyncflow/flowsvr/router"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	r := gin.Default()
	config.InitConfig()
	router.InitRouter(r)
	err := r.Run()
	if err != nil {
		log.Fatalf("run server failed, err: %s", err)
	}
}
