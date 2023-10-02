package user

import (
	"testing"
	"touchedFlowed/features/user"
)

func TestValidEmail(t *testing.T) {
	want := "qwikle@gmail.com"

	got, err := user.ValidEmail(want)
	if err != nil {
		t.Fatalf("ValidEmail(%q) = %q, want %q", want, got, want)
	}
}

func TestInvalidEmail(t *testing.T) {
	want := "qwikle@gmail"

	_, err := user.ValidEmail(want)
	if err == nil {
		t.Fatalf("ValidEmail(%q) = %q, want %q", want, err, want)
	}
}

func TestEmptyEmail(t *testing.T) {
	want := ""

	_, err := user.ValidEmail(want)
	if err == nil {
		t.Fatalf("ValidEmail(%q) = %q, want %q", want, err, want)
	}
}
