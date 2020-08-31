package github_provider

import (
	"errors"
	"github.com/TeplyyMaksim/golang-microservices/src/api/clients/rest_client"
	"github.com/TeplyyMaksim/golang-microservices/src/api/domain/github"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "Authorization", headerAuthorization)
	assert.EqualValues(t, "token %s", headerAuthorizationFormat)
	assert.EqualValues(t, "https://api.github.com/user/repos", urlCreateRepo)
}

func TestMain(m  *testing.M) {
	rest_client.StartMockups()
	os.Exit(m.Run())
}

func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("abc123")
	assert.EqualValues(t, header, "token abc123")
}

func TestCreateRepoErrorRestClient (t *testing.T) {
	rest_client.FlushMockups()
	rest_client.AddMockup(
		rest_client.Mock{
			Url:        "https://api.github.com/user/repos",
			HttpMethod: http.MethodPost,
			Error:      errors.New("invalid client response"),
		},
	)

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.StatusCode, http.StatusInternalServerError)
}

func TestCreateRepoErrorInvalidResponseBody (t *testing.T) {
	rest_client.FlushMockups()
	invalidCloser, _ := os.Open("-asf3")
	rest_client.AddMockup(
		rest_client.Mock{
			Url:        "https://api.github.com/user/repos",
			HttpMethod: http.MethodPost,
			Response: 	&http.Response{
				StatusCode:       http.StatusCreated,
				Body:             invalidCloser,
			},
		},
	)

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
}

func TestCreateRepoErrorInvalidErrorInterface (t *testing.T) {
	rest_client.FlushMockups()

	rest_client.AddMockup(
		rest_client.Mock{
			Url:        "https://api.github.com/user/repos",
			HttpMethod: http.MethodPost,
			Response: 	&http.Response{
				StatusCode:       http.StatusInternalServerError,
				Body:             ioutil.NopCloser(strings.NewReader(`{ "message": 1 }`)),
			},
		},
	)

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid json create repo error body", err.Message)
}

func TestCreateRepoUnauthorized (t *testing.T) {
	rest_client.FlushMockups()

	rest_client.AddMockup(
		rest_client.Mock{
			Url:        "https://api.github.com/user/repos",
			HttpMethod: http.MethodPost,
			Response: 	&http.Response{
				StatusCode:       http.StatusUnauthorized,
				Body:             ioutil.NopCloser(strings.NewReader(`{ "message": "Authentication required","document_url": "https://idk.com" }`)),
			},
		},
	)

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode)
	assert.EqualValues(t, "Authentication required", err.Message)
}

func TestCreateRepoInvalidSuccessResponse (t *testing.T) {
	rest_client.FlushMockups()

	rest_client.AddMockup(
		rest_client.Mock{
			Url:        "https://api.github.com/user/repos",
			HttpMethod: http.MethodPost,
			Response: 	&http.Response{
				StatusCode:       http.StatusCreated,
				Body:             ioutil.NopCloser(strings.NewReader(`{ "id": "123" }`)),
			},
		},
	)

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "invalid success create repo response body", err.Message)
}

func TestCreateRepoNoError (t *testing.T) {
	rest_client.FlushMockups()

	rest_client.AddMockup(
		rest_client.Mock{
			Url:        "https://api.github.com/user/repos",
			HttpMethod: http.MethodPost,
			Response: 	&http.Response{
				StatusCode:       http.StatusCreated,
				Body:             ioutil.NopCloser(strings.NewReader(`{ "id": 212412, "name": "golang_tutorial", "full_name": "Alvares" }`)),
			},
		},
	)

	response, err := CreateRepo("", github.CreateRepoRequest{})

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.EqualValues(t, response.Id, 212412)
}