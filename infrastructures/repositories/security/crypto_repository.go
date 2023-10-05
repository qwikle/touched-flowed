package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"touchedFlowed/features/utils"
)
import "golang.org/x/crypto/bcrypt"

type HashRepository struct {
	passwordHashes *utils.Hashes
}

const keySize = 32

func deriveKey(key string) []byte {
	hash := sha256.Sum256([]byte(key))
	return hash[:]
}

func generateNonce() ([]byte, error) {
	nonce := make([]byte, 12)
	_, err := rand.Read(nonce)
	if err != nil {
		return nil, err
	}
	return nonce, nil
}

func (b HashRepository) Encode(message interface{}) (string, error) {
	apiKey := os.Getenv("API_KEY")
	key := deriveKey(apiKey)

	plaintext, err := json.Marshal(message)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce, err := generateNonce()
	if err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)

	nonceBase64 := base64.StdEncoding.EncodeToString(nonce)
	messageBase64 := base64.StdEncoding.EncodeToString(ciphertext)

	result := messageBase64 + "." + nonceBase64

	return result, nil
}

func (b HashRepository) Decode(encoded string) (interface{}, error) {
	apiKey := os.Getenv("API_KEY")
	key := deriveKey(apiKey)

	parts := strings.Split(encoded, ".")
	if len(parts) != 2 {
		return nil, errors.New("invalid format")
	}

	messageBase64, nonceBase64 := parts[0], parts[1]

	nonce, err := base64.StdEncoding.DecodeString(nonceBase64)
	if err != nil {
		return nil, err
	}

	ciphertext, err := base64.StdEncoding.DecodeString(messageBase64)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	var result interface{}

	err = json.Unmarshal(plaintext, &result)
	return result, nil
}

func (b HashRepository) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (b HashRepository) Compare(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NewHashRepository() utils.Hashes {
	return &HashRepository{}
}
