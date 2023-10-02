package requests

import valueobjects "touchedFlowed/features/user/value-objects"

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ValidSignInRequest(u *SignInRequest) (*SignInRequest, error) {
	var err error
	var request SignInRequest

	request.Email, err = valueobjects.ValidEmail(u.Email)
	if err != nil {
		return nil, err
	}
	request.Password = u.Password
	return &request, nil
}
