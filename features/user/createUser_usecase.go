package user

import "errors"

type CreateUseCase interface {
	Execute(request *CreateUserRequest) (*CreateUserResponse, error)
}

type createUserUseCase struct {
	repository Repository
}

func (c createUserUseCase) Execute(request *CreateUserRequest) (*CreateUserResponse, error) {
	var response CreateUserResponse

	user, err := ValidCreateUserRequest(request)
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

func NewCreateUserUseCase(r Repository) CreateUseCase {
	return &createUserUseCase{
		repository: r,
	}
}
