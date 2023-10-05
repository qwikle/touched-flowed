package repository

import (
	"touchedFlowed/features/user/entities"
)

type TokenRepository interface {
	GenerateToken(token uint64) (string, error)
	DeleteToken(token string) error
	GetUserByToken(token string) (*entities.User, error)
}
