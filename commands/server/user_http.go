package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/finalapproachsoftware/leasing-api/api"

	kithttp "github.com/go-kit/kit/transport/http"
)

func (s *Server) registerUserRoutes() {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(jsonEncodeHttpError),
	}

	createUserHandler := kithttp.NewServer(
		s.api.UserCreateEndpoint,
		decodeCreateUserRequest,
		jsonEncodeHttpResponse,
		opts...,
	)

	s.httpRouter.Handle("/v1/users", createUserHandler).Methods("POST")
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request api.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
