package responses

import "touchedFlowed/features/user/entities"

type CreateUserResponse struct {
	Id        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (r *CreateUserResponse) FromEntity(user *entities.User) {
	r.Id = user.Id
	r.FirstName = user.FirstName
	r.LastName = user.LastName
	r.Email = user.Email
}
