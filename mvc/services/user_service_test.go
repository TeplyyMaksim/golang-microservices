package services

import (
	"github.com/TeplyyMaksim/golang-microservices/mvc/domain"
	"github.com/TeplyyMaksim/golang-microservices/mvc/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	userDaoTest userDaoMock
	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

type userDaoMock struct {}

func (m *userDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError)  {
	return getUserFunction(userId)
}

func init() {
	domain.UserDao = &userDaoMock{}
}

func TestUsersService_GetUserNotFound(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Message: "Not found",
			Status:  http.StatusNotFound,
			Code:    "not_found",
		}
	}

	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
}

func TestUserService_GetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			Id:        10,
			FirstName: "Maksym",
			LastName:  "Teplyy",
			Email:     "teplyy.maksim@whatever.com",
		}, nil
	}

	user, err := UserService.GetUser(10)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, user.Id, 10)
}