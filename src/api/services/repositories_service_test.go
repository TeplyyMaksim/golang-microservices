package services

import (
	"github.com/TeplyyMaksim/golang-microservices/src/api/clients/rest_client"
	"github.com/TeplyyMaksim/golang-microservices/src/api/domain/repositories"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	rest_client.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidInputName(t *testing.T) {
	request := repositories.CreateRepoRequest{}

	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status())
}

func TestCreateRepoErrorFromGithub(t *testing.T) {
	rest_client.FlushMockups()
	rest_client.AddMockup(rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response:   &http.Response{
			StatusCode: 	http.StatusUnauthorized,
			Body:       	ioutil.NopCloser(strings.NewReader(`{ "message": "Requires authentication", "documentation_url": "oh_no_no" }`)),
		},
	})

	request := repositories.CreateRepoRequest{ Name: "Asd-bsd-cpp" }
	result, err := RepositoryService.CreateRepo(request)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status())
	assert.EqualValues(t, "Requires authentication", err.Message())
}

func TestCreateRepoNoError(t *testing.T) {
	rest_client.FlushMockups()
	rest_client.AddMockup(rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response:   &http.Response{
			StatusCode: 	http.StatusCreated,
			Body:       	ioutil.NopCloser(strings.NewReader(`{ "id": 123 }`)),
		},
	})

	request := repositories.CreateRepoRequest{ Name: "Asd-bsd-cpp" }
	result, err := RepositoryService.CreateRepo(request)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, result.Id, 123)
}