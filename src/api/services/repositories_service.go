package services

import (
	"github.com/TeplyyMaksim/golang-microservices/src/api/config"
	"github.com/TeplyyMaksim/golang-microservices/src/api/domain/github"
	"github.com/TeplyyMaksim/golang-microservices/src/api/domain/repositories"
	"github.com/TeplyyMaksim/golang-microservices/src/api/providers/github_provider"
	"github.com/TeplyyMaksim/golang-microservices/src/api/utils/errors"
	"strings"
)

type repoService struct {

}

type repoServiceInterface interface {
	CreateRepo(repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

var (
	RepositoryService repoServiceInterface
)

func init()  {
	RepositoryService = &repoService{}
}

func (s *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)

	if input.Name == "" {
		return nil, errors.NewBedRequestError("Invalid repository name")
	}
	
	request := github.CreateRepoRequest{
		Name:        input.Name,
		Private:     true,
	}

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)

	if err != nil {
		apiErr := errors.NewApiError(err.StatusCode, err.Message)

		return nil, apiErr
	}

	result := repositories.CreateRepoResponse{
		Id: response.Id,
		Name: response.Name,
		Owner: response.Owner.Login,
	}

	return &result, nil
}