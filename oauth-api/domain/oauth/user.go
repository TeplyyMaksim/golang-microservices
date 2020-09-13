package oauth

import (
	"fmt"
	"github.com/TeplyyMaksim/golang-microservices/src/api/utils/errors"
)

type User struct {
	Id 			int64 	`json:"id"`
	Username 	string 	`json:"username"`
}

const (
	queryGetUserByUsernameAndPassword = "SELECT id, username FROM users WHERE username=? AND password=?;"
)

var (
	users = map[string]*User {
		"fede": &User{
			Id:       123,
			Username: "fede",
		},
	}
)

func GetUserByUsernameAndPassword(username string, password string) (*User, errors.ApiError) {
	user := users[username]

	if user == nil {
		return nil, errors.NewNotFoundApiError(
			fmt.Sprintf("no user found with given parameters"),
		)
	}

	return user, nil
}