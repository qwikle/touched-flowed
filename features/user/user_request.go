package user

type CreateUserRequest struct {
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Email             string `json:"email"`
	EmailConfirmation string `json:"email_confirmation"`
	Password          string `json:"password"`
	PassConfirmation  string `json:"password_confirmation"`
}

func ValidCreateUserRequest(u *CreateUserRequest) (*CreateUserRequest, error) {
	var err error
	var request CreateUserRequest

	request.FirstName, err = NameIsValid(u.FirstName, "first name")
	if err != nil {
		return nil, err
	}

	request.LastName, err = NameIsValid(u.LastName, "last name")
	if err != nil {
		return nil, err
	}

	request.Email, err = ValidEmail(u.Email)
	if err != nil {
		return nil, err
	}

	request.EmailConfirmation, err = EmailIsConfirmed(u.Email, u.EmailConfirmation)
	if err != nil {
		return nil, err
	}

	request.Password, err = PasswordIsValid(u.Password)
	if err != nil {
		return nil, err
	}

	request.PassConfirmation, err = PasswordIsConfirmed(u.Password, u.PassConfirmation)
	if err != nil {
		return nil, err
	}

	return &request, nil
}
