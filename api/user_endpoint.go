package api

import (
	"context"

	"github.com/finalapproachsoftware/leasing-api/model"
	"github.com/finalapproachsoftware/leasing-api/services"
)

type CreateUserRequest struct {
	FirstName string `json: "firstName'`
}

type CreateUserResponse struct {
	FirstName string `json: "firstName'`
	LastName  string `json: "lastName'`
}

func (api *Api) registerCreateUserEndpoint(svc services.IUserService) {
	api.UserCreateEndpoint = func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)

		user := model.User{
			FirstName: req.FirstName,
			LastName: "Unknown",
		}

		user, err := svc.Create(user)
		if err != nil {
			return nil, err
		}

		return CreateUserResponse{
			FirstName: user.FirstName,
			LastName: user.LastName,
		}, nil

	}
}
