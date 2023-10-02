package user

import (
	"touchedFlowed/features/utils"
)

type Repository interface {
	CreateUser(user *CreateUserRequest) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserById(id uint64) (*User, error)
	updateUser(user *User) (*User, error)
}

type repository struct {
	db utils.Database
}

func (r repository) CreateUser(user *CreateUserRequest) (*User, error) {
	row, err := r.db.Query("SELECT * FROM insert_user_json($1)", user.ToJson())
	if err != nil {
		return nil, err
	}
	var id uint64
	for row.Next() {
		err := row.Scan(&id)
		if err != nil {
			return nil, err
		}
	}

	return &User{
		Id:        id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}, nil
}

func (r repository) GetUserByEmail(email string) (*User, error) {
	result := r.db.QueryRow("SELECT * FROM users WHERE email = $1", email)

	var user User
	err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r repository) GetUserById(id uint64) (*User, error) {
	result := r.db.QueryRow("SELECT * FROM users WHERE id = $1", id)

	var user User
	err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r repository) updateUser(user *User) (*User, error) {
	exec, err := r.db.Exec("UPDATE users SET first_name = $1, last_name = $2, email = $3, password = $4 WHERE id = $5", user.FirstName, user.LastName, user.Email, user.Password, user.Id)
	if err != nil {
		return nil, err
	}
	id, err := exec.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &User{
		Id:        uint64(id),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}, nil
}

func NewRepository(db utils.Database) Repository {
	return &repository{
		db: db,
	}
}
