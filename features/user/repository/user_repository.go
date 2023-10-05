package repository

import (
	"touchedFlowed/features/user/entities"
	"touchedFlowed/features/user/requests"
)

type UserRepository interface {
	CreateUser(user *requests.CreateUserRequest) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	GetUserById(id uint64) (*entities.User, error)
	UpdateUser(user *entities.User) (*entities.User, error)
	DeleteUser(id uint64) error
	GetAll() ([]*entities.User, error)
}
