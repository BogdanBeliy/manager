package repository

import (
	"github.com/BogdanBeliy/manager/internal/models/entities"
	"github.com/BogdanBeliy/manager/internal/models/entities/requests"
)

type UserInterface interface {
	Create(request requests.CreateUserRequest)
}

type UserRepository struct {
	storage map[string]entities.User
}

func NewUserRepository() *UserRepository {
	r := UserRepository{
		storage: make(map[string]entities.User),
	}
	return &r
}
