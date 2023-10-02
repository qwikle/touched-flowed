package user

import "touchedFlowed/features/user/entities"
import "golang.org/x/crypto/bcrypt"

type BcryptPasswordHashes struct {
	passwordHashes *entities.PasswordHashes
}

func (b BcryptPasswordHashes) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (b BcryptPasswordHashes) Compare(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NewBcryptPasswordHashes() *BcryptPasswordHashes {
	return &BcryptPasswordHashes{}
}
