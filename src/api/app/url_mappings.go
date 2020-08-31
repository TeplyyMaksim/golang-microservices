package app

import (
	"github.com/TeplyyMaksim/golang-microservices/mvc/controllers/polo"
	"github.com/TeplyyMaksim/golang-microservices/src/api/controllers/repositories"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)
	router.POST("/repositories", repositories.CreateRepo)
}