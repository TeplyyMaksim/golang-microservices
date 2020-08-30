package controllers

import (
	"github.com/TeplyyMaksim/golang-microservices/mvc/services"
	"github.com/TeplyyMaksim/golang-microservices/mvc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context)  {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "user_id must be a number",
			Status:  http.StatusBadRequest,
			Code:    "bad_request",
		}

		utils.RespondError(c, apiErr)

		return
	}

	user, apiErr := services.UserService.GetUser(userId)

	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, user)
}