package oauth

import "github.com/TeplyyMaksim/golang-microservices/src/api/utils/errors"

var (
	tokens = make(map[string] *AccessToken, 0)
)

func (at *AccessToken) Save() errors.ApiError {
	at.AccessToken = "someRandomId"

	tokens[at.AccessToken] = at



	return nil
}

func GetAccessTokenByToken(accessToken string) (*AccessToken, *errors.ApiError) {
	token := tokens[accessToken]

	if token == nil {
		err :=  errors.NewNotFoundApiError("no access token found with given parameters")
		return nil, &err
	}

	return token, nil
}