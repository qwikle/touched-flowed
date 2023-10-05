package utils

type Hashes interface {
	Hash(password string) (string, error)
	Compare(password string, hash string) bool
	Encode(message interface{}) (string, error)
	Decode(encoded string) (interface{}, error)
}
