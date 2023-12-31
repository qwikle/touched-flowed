package usecases

import (
	"errors"
	"touchedFlowed/features/user/repository"
	"touchedFlowed/features/user/requests"
	"touchedFlowed/features/user/responses"
	"touchedFlowed/features/utils"
)

type CreateUseCase interface {
	Execute(request *requests.CreateUserRequest) (*responses.CreateUserResponse, error)
}

type createUserUseCase struct {
	repository repository.UserRepository
	hash       utils.Hashes
}

func (c createUserUseCase) Execute(request *requests.CreateUserRequest) (*responses.CreateUserResponse, error) {
	var response responses.CreateUserResponse

	user, err := requests.ValidCreateUserRequest(request)
	if err != nil {
		return nil, err
	}

	email, _ := c.repository.GetUserByEmail(user.Email)
	if email != nil {
		return nil, errors.New("email already exists")
	}

	user.Password, err = c.hash.Hash(user.Password)
	if err != nil {
		return nil, err
	}

	newUser, err := c.repository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	response.FromEntity(newUser)

	return &response, nil
}

func NewCreateUserUseCase(r repository.UserRepository, h utils.Hashes) CreateUseCase {
	return &createUserUseCase{
		repository: r,
		hash:       h,
	}
}
