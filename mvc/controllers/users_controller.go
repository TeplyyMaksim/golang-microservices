package controllers

import (
	"encoding/json"
	"github.com/TeplyyMaksim/golang-microservices/mvc/services"
	"github.com/TeplyyMaksim/golang-microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request)  {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "user_id must be a number",
			Status:  http.StatusBadRequest,
			Code:    "bad_request",
		}

		utils.HandleApplicationError(apiErr, resp)

		return
	}

	user, apiErr := services.UserService.GetUser(userId)

	if apiErr != nil {
		utils.HandleApplicationError(apiErr, resp)
		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}