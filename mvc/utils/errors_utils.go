package utils

import (
	"encoding/json"
	"net/http"
)

type ApplicationError struct {
	Message string 	`json:"message"`
	Status int 		`json:"status"`
	Code string 	`json:"code"`
}

func HandleApplicationError (applicationError *ApplicationError, resp http.ResponseWriter) {
	jsonValue, jsonError := json.Marshal(applicationError)

	if jsonError != nil {
		HandleApplicationError(
			&ApplicationError{
				Message: "Internal server error",
				Status:  http.StatusInternalServerError,
				Code:    "server_error",
			},
			resp,
		)

		return
	}

	resp.WriteHeader(applicationError.Status)
	resp.Write(jsonValue)
}