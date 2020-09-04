package github_provider

import (
	"encoding/json"
	"fmt"
	"github.com/TeplyyMaksim/golang-microservices/src/api/clients/rest_client"
	"github.com/TeplyyMaksim/golang-microservices/src/api/domain/github"
	"io/ioutil"
	"net/http"
)

const (
	headerAuthorization = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateRepo = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.ErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))
	response, err := rest_client.Post(urlCreateRepo, request, headers)
	fmt.Println("response", response)
	// Error while sending request
	if err != nil {
		return nil, &github.ErrorResponse{
			Message: "error while sending create repo request",
			StatusCode: http.StatusInternalServerError,
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	// Error while parsing response
	if err != nil {
		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message: "invalid response body",
		}
	}

	if response.StatusCode > 299 {
		var errResponse github.ErrorResponse

		// Error while parsing error response
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message: "invalid json create repo error body",
			}
		}

		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result github.CreateRepoResponse
	// Error while parsing successful
	if err := json.Unmarshal(bytes, &result);  err != nil {
		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message: "invalid success create repo response body",
		}
	}

	return &result, nil
}
