package user

import (
	"touchedFlowed/features/user/entities"
	"touchedFlowed/features/user/repository"
	"touchedFlowed/features/user/requests"
	"touchedFlowed/features/utils"
)

type pgUserRepository struct {
	db             utils.Database
	userRepository *repository.UserRepository
}

func scanAndReturnUser(r utils.Rows) (*entities.User, error) {
	var user entities.User
	err := r.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r pgUserRepository) GetAll() ([]*entities.User, error) {
	row, err := r.db.Query(`SELECT * FROM "users"`)
	if err != nil {
		return nil, err
	}
	var users []*entities.User
	for row.Next() {
		var user *entities.User
		user, err = scanAndReturnUser(row)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r pgUserRepository) DeleteUser(id uint64) error {
	_, err := r.db.Query(`DELETE FROM "users" WHERE id=$1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (r pgUserRepository) CreateUser(user *requests.CreateUserRequest) (*entities.User, error) {
	row, err := r.db.Query("SELECT * FROM insert_user_json($1)", user)
	if err != nil {
		return nil, err
	}
	var id uint64

	for row.Next() {
		err = row.Scan(&id)
		if err != nil {
			return nil, err
		}
	}

	newUser := &entities.User{
		Id:        id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}

	return newUser, nil
}

func (r pgUserRepository) GetUserByEmail(email string) (*entities.User, error) {
	row, err := r.db.Query(`SELECT * FROM "users" WHERE "email"= $1`, email)
	if err != nil {
		return nil, err
	}
	var user *entities.User
	for row.Next() {
		user, err = scanAndReturnUser(row)
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}

func (r pgUserRepository) GetUserById(id uint64) (*entities.User, error) {
	row, err := r.db.Query(`SELECT * FROM "users" WHERE "id"=$1`, id)
	if err != nil {
		return nil, err
	}
	var user *entities.User
	for row.Next() {
		user, err = scanAndReturnUser(row)
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}

func (r pgUserRepository) UpdateUser(user *entities.User) (*entities.User, error) {
	row, err := r.db.Query(`UPDATE "users" SET "first_name"=$1, "last_name"=$2, "email"=$3, "password"=$4 WHERE "id"=$5 RETURNING *`, user.FirstName, user.LastName, user.Email, user.Password, user.Id)
	if err != nil {
		return nil, err
	}
	var updatedUser *entities.User
	for row.Next() {
		updatedUser, err = scanAndReturnUser(row)
		if err != nil {
			return nil, err
		}
	}
	return updatedUser, nil
}

func NewPgUserRepository(db utils.Database) repository.UserRepository {
	return &pgUserRepository{
		db: db,
	}
}
