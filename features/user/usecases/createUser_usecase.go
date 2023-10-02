package usecases

import (
	"errors"
	"touchedFlowed/features/user/repository"
	"touchedFlowed/features/user/requests"
	"touchedFlowed/features/user/responses"
)

type CreateUseCase interface {
	Execute(request *requests.CreateUserRequest) (*responses.CreateUserResponse, error)
}

type createUserUseCase struct {
	repository repository.Repository
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

	newUser, err := c.repository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	response.Id = newUser.Id
	response.FirstName = newUser.FirstName
	response.LastName = newUser.LastName
	response.Email = newUser.Email

	return &response, nil
}

func NewCreateUserUseCase(r repository.Repository) CreateUseCase {
	return &createUserUseCase{
		repository: r,
	}
}
