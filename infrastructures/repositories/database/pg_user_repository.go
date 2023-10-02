package database

import (
	"touchedFlowed/features/user/entities"
	"touchedFlowed/features/user/repository"
	"touchedFlowed/features/user/requests"
	"touchedFlowed/features/utils"
)

type pgUserRepository struct {
	db             utils.Database
	userRepository *repository.Repository
}

func (r pgUserRepository) DeleteUser(id uint64) error {
	_, err := r.db.Query(`DELETE FROM "users" WHERE id=$1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (r pgUserRepository) CreateUser(user *requests.CreateUserRequest) (*entities.User, error) {
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

	return &entities.User{
		Id:        id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}, nil
}

func (r pgUserRepository) GetUserByEmail(email string) (*entities.User, error) {
	result := r.db.QueryRow("SELECT * FROM users WHERE email = $1", email)

	var user entities.User
	err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r pgUserRepository) GetUserById(id uint64) (*entities.User, error) {
	result := r.db.QueryRow("SELECT * FROM users WHERE id = $1", id)

	var user entities.User
	err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r pgUserRepository) UpdateUser(user *entities.User) (*entities.User, error) {
	exec, err := r.db.Exec("UPDATE users SET first_name = $1, last_name = $2, email = $3, password = $4 WHERE id = $5", user.FirstName, user.LastName, user.Email, user.Password, user.Id)
	if err != nil {
		return nil, err
	}
	id, err := exec.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &entities.User{
		Id:        uint64(id),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}, nil
}

func NewPgUserRepository(db utils.Database) repository.Repository {
	return &pgUserRepository{
		db: db,
	}
}
