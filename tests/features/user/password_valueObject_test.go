package user

import (
	"testing"
	"touchedFlowed/features/user"
)

func TestValidPassword(t *testing.T) {
	want := "V€ry$tr0ngP@ssw0rd"

	got, err := user.PasswordIsValid(want)
	if err != nil {
		t.Fatalf("ValidPassword(%q) = %q, want %q", want, got, want)
	}
}

func TestInvalidPassword(t *testing.T) {
	want := "weak"

	_, err := user.PasswordIsValid(want)
	if err == nil {
		t.Fatalf("ValidPassword(%q) = %q, want %q", want, err, want)
	}
}

func TestEmptyPassword(t *testing.T) {
	want := ""

	_, err := user.PasswordIsValid(want)
	if err == nil {
		t.Fatalf("ValidPassword(%q) = %q, want %q", want, err, want)
	}
}
