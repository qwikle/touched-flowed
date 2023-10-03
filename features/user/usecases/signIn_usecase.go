package usecases

import (
	"errors"
	"touchedFlowed/features/user/entities"
	"touchedFlowed/features/user/repository"
	"touchedFlowed/features/user/requests"
	"touchedFlowed/features/user/responses"
)

type SignInUseCase interface {
	Execute(request *requests.SignInRequest) (*responses.CreateUserResponse, error)
}

type signInUseCase struct {
	repository repository.Repository
	hash       entities.PasswordHashes
}

func (c signInUseCase) Execute(request *requests.SignInRequest) (*responses.CreateUserResponse, error) {
	var response responses.CreateUserResponse

	user, err := requests.ValidSignInRequest(request)
	if err != nil {
		return nil, err
	}

	newUser, _ := c.repository.GetUserByEmail(user.Email)

	if newUser == nil {
		return nil, errors.New("invalid email or password")
	}

	if !c.hash.Compare(user.Password, newUser.Password) {
		return nil, errors.New("invalid email or password")
	}

	response.FromEntity(newUser)

	return &response, nil
}

func NewSignInUseCase(repository repository.Repository, hash entities.PasswordHashes) SignInUseCase {
	return &signInUseCase{
		repository: repository,
		hash:       hash,
	}
}
