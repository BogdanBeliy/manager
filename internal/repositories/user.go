package repositories

import (
	"github.com/BogdanBeliy/manager/internal/domain/models/requests"
	"github.com/BogdanBeliy/manager/internal/domain/models/responses"
	"github.com/BogdanBeliy/manager/internal/domain/storage"
)

type UserInterface interface {
	Create(data requests.CreateUserRequest) (responses.CreateUserResponse, error)
	Update()
	Delete()
	FetchAll()
	FetchOne()
}

type UserRepository struct {
	storage map[string]storage.User
}

func NewUserRepository(storage map[string]storage.User) *UserRepository {
	r := UserRepository{
		storage: storage,
	}
	return &r
}

func (u *UserRepository) Create(data storage.User) (storage.User, error) {
	u.storage[data.Email] = data
	return u.storage[data.Email], nil
}
