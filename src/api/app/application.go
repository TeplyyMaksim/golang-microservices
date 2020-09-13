package app

import (
	"github.com/TeplyyMaksim/golang-microservices/src/api/log/option_a"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	option_a.Info("about to map the urls", "step:1", "status:pending")
	mapUrls()
	option_a.Info("urls successfully mapped", "step:2", "status:success")

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}