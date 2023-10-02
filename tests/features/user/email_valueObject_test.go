package user

import (
	"testing"
	"touchedFlowed/features/user"
)

func testEmail(t *testing.T) string {
	email, err := user.ValidEmail("qwikle@gmail.com")
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	return email

}
