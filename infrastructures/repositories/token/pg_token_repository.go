package token

import (
	"crypto/subtle"
	"encoding/json"
	"errors"
	"strings"
	"time"
	"touchedFlowed/features/user/entities"
	"touchedFlowed/features/user/repository"
	"touchedFlowed/features/utils"
)

type PgTokenRepository struct {
	db             utils.Database
	hash           utils.Hashes
	memory         utils.MemoryDatabase
	userRepository repository.UserRepository
}

func (p PgTokenRepository) DeleteToken(token string) error {
	err := p.memory.Delete(token)
	if err != nil {
		return err
	}
	return nil
}

func (p PgTokenRepository) GetUserByToken(token string) (*entities.User, error) {
	decodedToken, err := p.GetToken(token)
	if err != nil {
		return nil, err
	}
	user, err := p.userRepository.GetUserById(decodedToken.Id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (p PgTokenRepository) GenerateToken(id uint64) (string, error) {
	encodedKey, err := p.hash.Encode(id)
	if err != nil {
		return "", err
	}
	token := &entities.Token{Id: id}
	encoded, err := p.hash.Encode(token)
	if err != nil {
		return "", err
	}
	token.Token = encoded
	err = p.StoreToken(token, encodedKey)
	if err != nil {
		return "", err
	}
	return encoded + ":" + encodedKey, nil
}

func (p PgTokenRepository) StoreToken(token *entities.Token, hashedToken string) error {
	stringedToken, err := json.Marshal(token)
	if err != nil {
		return err
	}
	key := "api:" + hashedToken
	err = p.memory.SetWithExpiration(key, string(stringedToken), 72*time.Hour)
	if err != nil {
		return err
	}
	return nil
}

func (p PgTokenRepository) CheckToken(token string, token2 string) bool {
	secretToken := []byte(token)
	secretToken2 := []byte(token2)
	if len(secretToken) != len(secretToken2) {
		return false
	}
	return subtle.ConstantTimeCompare(secretToken, secretToken2) == 1
}

func (p PgTokenRepository) GetToken(hashedToken string) (*entities.Token, error) {
	token, encodedKey := strings.Split(hashedToken, ":")[0], strings.Split(hashedToken, ":")[1]
	if token == "" || encodedKey == "" {
		return nil, errors.New("invalid token")
	}
	existedToken, err := p.memory.Get("api:" + encodedKey)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	if existedToken == "" {
		return nil, errors.New("invalid token")
	}

	var decodedToken entities.Token
	err = json.Unmarshal([]byte(existedToken), &decodedToken)
	if err != nil {
		return nil, err
	}
	if !p.CheckToken(token, decodedToken.Token) {
		return nil, errors.New("invalid token")
	}
	return &decodedToken, nil
}

func NewPgTokenRepository(memory utils.MemoryDatabase, userRepository repository.UserRepository, db utils.Database, hash utils.Hashes) repository.TokenRepository {
	return &PgTokenRepository{db: db, hash: hash, memory: memory, userRepository: userRepository}
}
