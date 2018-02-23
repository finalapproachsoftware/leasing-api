package api

import (
	"github.com/finalapproachsoftware/leasing-api/services"
	"github.com/go-kit/kit/endpoint"
)

type Api struct {
	UserCreateEndpoint endpoint.Endpoint
}

func NewApi() (*Api, error) {
	api := &Api{}

	userSvc := &services.UserService{}
	api.registerCreateUserEndpoint(userSvc)

	return api, nil
}
