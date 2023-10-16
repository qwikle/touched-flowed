package repository

import (
	"errors"
	"touchedFlowed/features/user/entities"
	"touchedFlowed/features/user/repository"
	"touchedFlowed/features/user/requests"
)

type UserMockRepository struct {
}

func (u UserMockRepository) CreateUser(user *requests.CreateUserRequest) (*entities.User, error) {
	validUser, err := requests.ValidCreateUserRequest(user)
	if err != nil {
		return nil, err
	}
	return &entities.User{
		Id:        1,
		FirstName: validUser.FirstName,
		LastName:  validUser.LastName,
		Email:     validUser.Email,
		Password:  validUser.Password,
	}, nil

}

func (u UserMockRepository) GetUserByEmail(email string) (*entities.User, error) {
	if email == "qwikle@gmail.com" {
		return &entities.User{
			Id:        1,
			FirstName: "Qwikle",
			LastName:  "User",
			Email:     "qwikle@gmail.com",
			Password:  "V€ry$tr0ngP@ssw0rd",
		}, nil
	}
	return nil, errors.New("user not found")
}

func (u UserMockRepository) GetUserById(id uint64) (*entities.User, error) {
	if id == 1 {
		return &entities.User{
			Id:        1,
			FirstName: "Qwikle",
			LastName:  "User",
			Email:     "qwikle@gmail.com",
			Password:  "V€ry$tr0ngP@ssw0rd",
		}, nil
	}
	return nil, errors.New("user not found")
}

func (u UserMockRepository) UpdateUser(user *entities.User) (*entities.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserMockRepository) DeleteUser(id uint64) error {
	if id == 1 {
		return nil
	}
	return errors.New("user not found")
}

func (u UserMockRepository) GetAll() ([]*entities.User, error) {
	return []*entities.User{
		{Id: 1,
			FirstName: "Qwikle",
			LastName:  "User",
			Email:     "qwikle@gmail.com",
			Password:  "V€ry$tr0ngP@ssw0rd"},
		{Id: 2,
			FirstName: "Caline",
			LastName:  "Bruno",
			Email:     "caline@gmail.com",
			Password:  "V€ry$tr0ngP@ssw0rd"},
	}, nil
}

func NewUserMockRepository() repository.UserRepository {
	return &UserMockRepository{}
}
