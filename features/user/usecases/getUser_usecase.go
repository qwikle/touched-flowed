package usecases

import (
	"touchedFlowed/features/user/repository"
	"touchedFlowed/features/user/responses"
)

type GetUserUseCase interface {
	Execute(token string) (*responses.CreateUserResponse, error)
}

type getUserUseCase struct {
	tokenRepository repository.TokenRepository
	userRepository  repository.UserRepository
}

func (g getUserUseCase) Execute(token string) (*responses.CreateUserResponse, error) {
	user, err := g.tokenRepository.GetUserByToken(token)
	if err != nil {
		return nil, err
	}
	var response responses.CreateUserResponse
	response.FromEntity(user)
	return &response, nil
}

func NewGetUserUseCase(tokenRepository *repository.TokenRepository, userRepository *repository.UserRepository) GetUserUseCase {
	return &getUserUseCase{tokenRepository: *tokenRepository, userRepository: *userRepository}
}
