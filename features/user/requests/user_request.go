package requests

import "touchedFlowed/features/user/value-objects"

type CreateUserRequest struct {
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Email             string `json:"email"`
	EmailConfirmation string `json:"email_confirmation"`
	Password          string `json:"password"`
	PassConfirmation  string `json:"password_confirmation"`
}

func (u *CreateUserRequest) ToJson() string {
	return `{"first_name": "` + u.FirstName + `", "last_name": "` + u.LastName + `", "email": "` + u.Email + `", "password": "` + u.Password + `"}`
}

func ValidCreateUserRequest(u *CreateUserRequest) (*CreateUserRequest, error) {
	var err error
	var request CreateUserRequest

	request.FirstName, err = value_objects.NameIsValid(u.FirstName, "first name")
	if err != nil {
		return nil, err
	}

	request.LastName, err = value_objects.NameIsValid(u.LastName, "last name")
	if err != nil {
		return nil, err
	}

	request.Email, err = value_objects.ValidEmail(u.Email)
	if err != nil {
		return nil, err
	}

	request.EmailConfirmation, err = value_objects.EmailIsConfirmed(u.Email, u.EmailConfirmation)
	if err != nil {
		return nil, err
	}

	request.Password, err = value_objects.PasswordIsValid(u.Password)
	if err != nil {
		return nil, err
	}

	request.PassConfirmation, err = value_objects.PasswordIsConfirmed(u.Password, u.PassConfirmation)
	if err != nil {
		return nil, err
	}

	return &request, nil
}
