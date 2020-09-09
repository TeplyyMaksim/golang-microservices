package repositories

import (
	"github.com/TeplyyMaksim/golang-microservices/src/api/utils/errors"
	"strings"
)

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r *CreateRepoRequest) Validate() errors.ApiError {
	r.Name = strings.TrimSpace(r.Name)

	if r.Name == "" {
		return errors.NewBadRequestError("Invalid repository name")
	}

	return nil
}

type CreateRepoResponse struct {
	Id			int64 		`json:"id"`
	Name		string		`json:"name"`
	Owner		string		`json:"owner"`
}

type CreateReposResponse struct {
	StatusCode 	int 							`json:"status"`
	Results     	[]CreateRepositoriesResult 		`json:"result"`
}

type CreateRepositoriesResult struct {
	Response 	*CreateRepoResponse 	`json:"response"`
	Error		errors.ApiError		`json:"error"`
}