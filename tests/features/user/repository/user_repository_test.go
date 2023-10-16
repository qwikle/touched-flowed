package repository

import (
	"testing"
	"touchedFlowed/features/user/requests"
)

func TestValidCreateUser(t *testing.T) {
	mock := NewUserMockRepository()

	user, err := mock.CreateUser(&requests.CreateUserRequest{
		FirstName:         "Qwikle",
		LastName:          "User",
		Email:             "qwikle@gmail.com",
		Password:          "V€ry$tr0ngP@ssw0rd",
		EmailConfirmation: "qwikle@gmail.com",
		PassConfirmation:  "V€ry$tr0ngP@ssw0rd",
	})
	if err != nil {
		t.Fatalf("CreateUser() = %q, want %q", err, user)
	}
}

func TestInvalidCreateUser(t *testing.T) {
	mock := NewUserMockRepository()

	user, _ := mock.CreateUser(&requests.CreateUserRequest{
		FirstName:         "Qwikle",
		LastName:          "User",
		Email:             "qwikle@gmail.com",
		Password:          "V€ry$tr0ngP@ssw0rd",
		EmailConfirmation: "qwikle@gmail.com",
	})
	if user != nil {
		t.Fatal("user should be nil")
	}
}

func TestGetUserByEmail(t *testing.T) {
	mock := NewUserMockRepository()

	user, err := mock.GetUserByEmail("qwikle@gmail.com")
	if err != nil {
		t.Fatalf("GetUserByEmail() = %q, want %q", err, user)
	}
}

func TestGetUserById(t *testing.T) {
	mock := NewUserMockRepository()

	user, err := mock.GetUserById(1)
	if err != nil {
		t.Fatalf("GetUserById() = %q, want %q", err, user)
	}
}
