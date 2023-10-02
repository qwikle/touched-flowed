package value_objects

import (
	"errors"
	"regexp"
)

func ValidEmail(email string) (string, error) {
	isValid := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString(email)
	if !isValid {
		return "", errors.New("invalid email")
	}
	return email, nil
}

func EmailIsConfirmed(email, confirmation string) (string, error) {
	if email != confirmation {
		return "", errors.New("email and confirmation must match")
	}
	return confirmation, nil
}
