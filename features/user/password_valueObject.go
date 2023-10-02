package user

import (
	"fmt"
	"regexp"
)

func PasswordIsValid(password string) (string, error) {
	// password must be at least 12 characters long and contain at least one uppercase letter, one lowercase letter, one number, and one special character (!?@&²*.;:), regex should be a valid go regex
	isValid := regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()_+\-=\[\]{};':"\\|,.<>/?]{12,}`).MatchString(password)
	if !isValid {
		return "", fmt.Errorf("password must be at least 12 characters long and contain at least one uppercase letter, one lowercase letter, one number, and one special character")
	}
	return password, nil
}

func PasswordIsConfirmed(password, confirmation string) (string, error) {
	if password != confirmation {
		return "", fmt.Errorf("password and confirmation must match")
	}
	return confirmation, nil
}
