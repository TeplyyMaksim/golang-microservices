package services

import (
	"github.com/TeplyyMaksim/golang-microservices/src/api/clients/rest_client"
	"github.com/TeplyyMaksim/golang-microservices/src/api/domain/repositories"
	"github.com/TeplyyMaksim/golang-microservices/src/api/utils/errors"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
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

func TestCreateRepoConcurrentInvalidRequest(t *testing.T) {
	request := repositories.CreateRepoRequest{}
	output := make(chan repositories.CreateRepositoriesResult)

	service := repoService{}
	go service.createRepoConcurrent(request, output)

	result := <- output
	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, http.StatusBadRequest, result.Error.Status())
}

func TestCreateRepoConcurrentGithubError(t *testing.T) {
	rest_client.FlushMockups()
	rest_client.AddMockup(rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response:   &http.Response{
			StatusCode: 	http.StatusUnauthorized,
			Body:       	ioutil.NopCloser(strings.NewReader(`{ "message": "Requires authentication", "documentation_url": "oh_no_no" }`)),
		},
	})
	request := repositories.CreateRepoRequest{Name: "testing"}
	output := make(chan repositories.CreateRepositoriesResult)

	service := repoService{}
	go service.createRepoConcurrent(request, output)

	result := <- output
	assert.NotNil(t, result)
	assert.Nil(t, result.Response)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, http.StatusUnauthorized, result.Error.Status())
}

func TestCreateRepoConcurrentNoError(t *testing.T) {
	rest_client.FlushMockups()
	rest_client.AddMockup(rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response:   &http.Response{
			StatusCode: 	http.StatusCreated,
			Body:       	ioutil.NopCloser(strings.NewReader(`{ "id": 123 }`)),
		},
	})
	request := repositories.CreateRepoRequest{Name: "testing"}
	output := make(chan repositories.CreateRepositoriesResult)

	service := repoService{}
	go service.createRepoConcurrent(request, output)

	result := <- output
	assert.NotNil(t, result)
	assert.NotNil(t, result.Response)
	assert.Nil(t, result.Error)
	assert.EqualValues(t, result.Response.Id, 123)
}

func TestHandleRepoResults (t *testing.T) {
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	var wg sync.WaitGroup

	service := repoService{}
	go service.handleRepoResults(&wg, input, output)

	wg.Add(1)
	go func () {
		input <- repositories.CreateRepositoriesResult{
			Error: errors.NewBadRequestError("Invalid repository name"),
		}
	}()

	wg.Wait()
	close(input)

	result := <- output

	assert.NotNil(t, result)
	assert.EqualValues(t, 0, result.StatusCode)
	assert.EqualValues(t, 1, len(result.Results))
	assert.EqualValues(t, http.StatusBadRequest, result.Results[0].Error.Status())
}

func TestCreateReposInvalidRequest(t *testing.T) {
	requests := []repositories.CreateRepoRequest{
		{},
		{ Name: "    " },
	}

	result, err := RepositoryService.CreateRepos(requests)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, 2, len(result.Results))
	assert.EqualValues(t, http.StatusBadRequest, result.StatusCode)

	assert.Nil(t, result.Results[0].Response)
	assert.EqualValues(t, http.StatusBadRequest, result.Results[0].Error.Status())

	assert.Nil(t, result.Results[1].Response)
	assert.EqualValues(t, http.StatusBadRequest, result.Results[1].Error.Status())
}

func TestCreateReposOneSuccessOneFail(t *testing.T) {
	rest_client.FlushMockups()
	rest_client.AddMockup(rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response:   &http.Response{
			StatusCode: 	http.StatusCreated,
			Body:       	ioutil.NopCloser(strings.NewReader(`{ "id": 123 }`)),
		},
	})

	requests := []repositories.CreateRepoRequest{
		{},
		{ Name: "testing" },
	}

	result, err := RepositoryService.CreateRepos(requests)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, 2, len(result.Results))
	assert.EqualValues(t, http.StatusPartialContent, result.StatusCode)

	for _, result := range result.Results {
		if result.Error != nil {
			assert.EqualValues(t, http.StatusBadRequest, result.Error.Status())
			continue
		}

		assert.EqualValues(t, result.Response.Id, 123)
	}

}

func TestCreateReposAllSuccess(t *testing.T) {
	rest_client.FlushMockups()
	rest_client.AddMockup(rest_client.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response:   &http.Response{
			StatusCode: 	http.StatusCreated,
			Body:       	ioutil.NopCloser(strings.NewReader(`{ "id": 123 }`)),
		},
	})

	requests := []repositories.CreateRepoRequest{
		{ Name: "testing" },
		{ Name: "stepping" },
	}

	result, err := RepositoryService.CreateRepos(requests)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, 2, len(result.Results))
	assert.EqualValues(t, http.StatusCreated, result.StatusCode)

	for _, result := range result.Results {
		assert.Nil(t, result.Error)
		assert.EqualValues(t, result.Response.Id, 123)
	}

}