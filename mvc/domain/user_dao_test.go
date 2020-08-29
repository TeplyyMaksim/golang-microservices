package domain

import (
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := UserDao.GetUser(0)

	assert.Nil(t, user,  "We don't expect user with id equals to 0")
	assert.NotNil(t, err, "We expect an error when id is 0")
	assert.EqualValues(t, err.Status, http.StatusNotFound, "We expect 404 status when user is not found")
	assert.EqualValues(t, err.Code, "not_found", "We expect not_found error code when user is not found")
}

func TestGetUserNotError(t *testing.T) {
	user, err := UserDao.GetUser(123)

	assert.Nil(t, err, "We expect no error with user id equals to 123")
	assert.NotNil(t, user, "We expect that 123 user exists")
	assert.EqualValues(t, 123, user.Id, "We expect that 123 user Id is 123")
}

