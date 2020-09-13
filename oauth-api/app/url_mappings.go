package app

import (
	"github.com/TeplyyMaksim/golang-microservices/oauth-api/controllers/oauth"
	"github.com/TeplyyMaksim/golang-microservices/src/api/controllers/polo"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)
}