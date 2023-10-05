package usecases

import (
	"errors"
	"touchedFlowed/features/user/repository"
	"touchedFlowed/features/user/requests"
	"touchedFlowed/features/user/responses"
	"touchedFlowed/features/utils"
)

type SignInUseCase interface {
	Execute(request *requests.SignInRequest) (*responses.SignInResponse, error)
}

type signInUseCase struct {
	userRepository  repository.UserRepository
	tokenRepository repository.TokenRepository
	hash            utils.Hashes
}

func (c signInUseCase) Execute(request *requests.SignInRequest) (*responses.SignInResponse, error) {
	user, err := requests.ValidSignInRequest(request)
	if err != nil {
		return nil, err
	}

	newUser, _ := c.userRepository.GetUserByEmail(user.Email)

	if newUser == nil {
		return nil, errors.New("invalid email or password")
	}

	if !c.hash.Compare(user.Password, newUser.Password) {
		return nil, errors.New("invalid email or password")
	}
	token, err := c.tokenRepository.GenerateToken(newUser.Id)
	if err != nil {
		return nil, err
	}

	return &responses.SignInResponse{
		Token: token,
	}, nil
}

func NewSignInUseCase(r repository.UserRepository, t repository.TokenRepository, hash utils.Hashes) SignInUseCase {
	return &signInUseCase{
		userRepository:  r,
		tokenRepository: t,
		hash:            hash,
	}
}
