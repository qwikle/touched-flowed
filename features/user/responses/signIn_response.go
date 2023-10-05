package responses

type SignInResponse struct {
	Token string `json:"token"`
}

func (r *SignInResponse) FromEntity(token string) {
	r.Token = token
}
