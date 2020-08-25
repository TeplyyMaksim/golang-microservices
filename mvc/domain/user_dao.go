package domain

import (
	"fmt"
	"github.com/TeplyyMaksim/golang-microservices/mvc/utils"
	"net/http"
)

var (
	users = map[int64] *User {
		123: &User { Id: 123, FirstName: "Maksym", LastName: "Teplyy", Email: "maksym@mailinator.com" },
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}

	//return nil, errors.New(fmt.Sprintf("User %v was not found\n", userId))
	return nil, &utils.ApplicationError{
		Message: fmt.Sprintf("User %v was not found\n", userId),
		Status: http.StatusNotFound,
		Code: "not_found",
	}
}