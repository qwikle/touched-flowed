package entities

type PasswordHashes interface {
	Hash(password string) (string, error)
	Compare(password string, hash string) bool
}
