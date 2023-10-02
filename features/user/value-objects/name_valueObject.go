package value_objects

import (
	"fmt"
	"regexp"
)

func NameIsValid(name, field string) (string, error) {
	isValid := regexp.MustCompile(`^[a-zA-Z]{2,}([ -][a-zA-Z]{2,})?$`).MatchString(name)
	if !isValid {
		return "", fmt.Errorf("invalid %s", field)
	}
	return name, nil
}
