package value_objects

import (
	"testing"
	"touchedFlowed/features/user/value-objects"
)

func TestValidPassword(t *testing.T) {
	want := "V€ry$tr0ngP@ssw0rd"

	got, err := value_objects.PasswordIsValid(want)
	if err != nil {
		t.Fatalf("ValidPassword(%q) = %q, want %q", want, got, want)
	}
}

func TestInvalidPassword(t *testing.T) {
	want := "weak"

	_, err := value_objects.PasswordIsValid(want)
	if err == nil {
		t.Fatalf("ValidPassword(%q) = %q, want %q", want, err, want)
	}
}

func TestEmptyPassword(t *testing.T) {
	want := ""

	_, err := value_objects.PasswordIsValid(want)
	if err == nil {
		t.Fatalf("ValidPassword(%q) = %q, want %q", want, err, want)
	}
}
