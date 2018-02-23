package services

import (
	"github.com/finalapproachsoftware/leasing-api/model"
)

type IUserService interface {
	Create(user model.User) (model.User, error)
}

type UserService struct {
}

func (s *UserService) Create(user model.User) (model.User, error) {
	return user, nil
}
