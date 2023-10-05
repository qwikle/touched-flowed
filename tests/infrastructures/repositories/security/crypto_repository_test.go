package security

import (
	"github.com/joho/godotenv"
	"testing"
	"touchedFlowed/features/user/entities"
	"touchedFlowed/infrastructures/repositories/security"
)

func TestEncode(t *testing.T) {
	godotenv.Load("../../../../.env")
	hashRepository := security.NewHashRepository()
	message := "Hello World"
	_, err := hashRepository.Encode(message)
	if err != nil {
		t.Error(err)
	}
}

func TestDecodeObject(t *testing.T) {
	hashRepository := security.NewHashRepository()
	token := entities.Token{
		Id: 456,
	}
	encoded, _ := hashRepository.Encode(token)
	decoded, err := hashRepository.Decode(encoded)

	decodedToken := entities.TokenFromJson(decoded)

	if decodedToken.Id != token.Id {
		t.Error("Decoded token is not equal to token")
	}
	if err != nil {
		t.Error(err)
	}
}

func TestDecodeString(t *testing.T) {
	hashRepository := security.NewHashRepository()
	encoded, _ := hashRepository.Encode("Hello World")
	_, err := hashRepository.Decode(encoded)
	if err != nil {
		t.Error(err)
	}
}
