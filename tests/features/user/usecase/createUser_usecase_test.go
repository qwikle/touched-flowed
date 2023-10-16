package usecase

import (
	"testing"
	"touchedFlowed/features/user/requests"
	"touchedFlowed/features/user/usecases"
	"touchedFlowed/tests/features/user/repository"
)

func TestCreateUserUseCase(t *testing.T) {
	mock := repository.NewUserMockRepository()
	user := &requests.CreateUserRequest{
		Email:             "qwikle1@gmail.com",
		EmailConfirmation: "qwikle1@gmail.com",
		FirstName:         "qwikle",
		LastName:          "qwikle",
		Password:          "V€ryS€cr€tP@ssw0rd",
		PassConfirmation:  "V€ryS€cr€tP@ssw0rd",
	}
	result, _ := usecases.NewCreateUserUseCase(mock, nil).Execute(user)

	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}
