package services

import (
	"github.com/TeplyyMaksim/golang-microservices/mvc/domain"
	"github.com/TeplyyMaksim/golang-microservices/mvc/utils"
	"net/http"
)

type itemsService struct {}

var (
	ItemsService itemsService
)

func (s *itemsService) GetItem(userId string) (*domain.Item, *utils.ApplicationError) {
	  return nil, &utils.ApplicationError{
	  	Message: "implement me",
	  	Status: http.StatusInternalServerError,
	  }
}