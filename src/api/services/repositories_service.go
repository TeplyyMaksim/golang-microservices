package services

import (
	"fmt"
	"github.com/TeplyyMaksim/golang-microservices/src/api/config"
	"github.com/TeplyyMaksim/golang-microservices/src/api/domain/github"
	"github.com/TeplyyMaksim/golang-microservices/src/api/domain/repositories"
	"github.com/TeplyyMaksim/golang-microservices/src/api/log/option_a"
	"github.com/TeplyyMaksim/golang-microservices/src/api/log/option_b"
	"github.com/TeplyyMaksim/golang-microservices/src/api/providers/github_provider"
	"github.com/TeplyyMaksim/golang-microservices/src/api/utils/errors"
	"net/http"
	"sync"
)

type repoService struct {

}

type repoServiceInterface interface {
	CreateRepo(clientId string, input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos([]repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
}

var (
	RepositoryService repoServiceInterface
)

func init()  {
	RepositoryService = &repoService{}
}

func (s *repoService) CreateRepo(clientId string, input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     true,
	}
	option_b.Info("about to send request to external api", fmt.Sprintf("client_id:%s", clientId), "status:pending")

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)

	if err != nil {
		option_b.Error("about to send request to external api", err, option_b.Field("client_id", clientId), "status:error")
		apiErr := errors.NewApiError(err.StatusCode, err.Message)

		return nil, apiErr
	}

	option_b.Info("response obtained from external api", option_b.Field("client_id", clientId), "status:success")
	result := repositories.CreateRepoResponse{
		Id: response.Id,
		Name: response.Name,
		Owner: response.Owner.Login,
	}

	return &result, nil
}

func (s *repoService) CreateRepos(requests []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError) {
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	defer close(output)

	var wg sync.WaitGroup
	go s.handleRepoResults(&wg, input, output)

	for _, current := range requests {
		wg.Add(1)
		go s.createRepoConcurrent(current, input)
	}

	wg.Wait()
	close(input)

	result := <- output

	successCreations := 0
	for _, current := range result.Results {
		if current.Response != nil {
			successCreations += 1
		}
	}


	if successCreations == 0 {
		result.StatusCode = result.Results[0].Error.Status()
	} else if successCreations == len(requests) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}

	return result, nil
}

func (s *repoService) handleRepoResults(wg *sync.WaitGroup,input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse) {
	var results repositories.CreateReposResponse

	for incomingEvent := range input {
		repoResult := repositories.CreateRepositoriesResult {
			Response: incomingEvent.Response,
			Error: incomingEvent.Error,
		}
		results.Results = append(results.Results, repoResult)

		wg.Done()
	}

	output <- results
}

func (s *repoService) createRepoConcurrent(input repositories.CreateRepoRequest, output chan repositories.CreateRepositoriesResult) {
	if err := input.Validate(); err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}

	result, err := s.CreateRepo(input)

	if err != nil {
		output <- repositories.CreateRepositoriesResult{Error:    err}
		return
	}
	output <- repositories.CreateRepositoriesResult{Response: result}
}