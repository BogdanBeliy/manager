package services

import (
	"github.com/BogdanBeliy/manager/internal/domain/models/requests"
	"github.com/BogdanBeliy/manager/internal/domain/storage"
	"github.com/BogdanBeliy/manager/internal/repositories"
)

type UserService struct {
	repository *repositories.UserRepository
}

func NewUserService(storage map[string]storage.User) *UserService {
	u := UserService{
		repository: repositories.NewUserRepository(storage),
	}
	return &u
}

func (u *UserService) Create(data requests.CreateUserRequest) ([]byte, error) {

	return nil, nil
}
