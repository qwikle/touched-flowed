package value_objects

import (
	"testing"
	"touchedFlowed/features/user/value-objects"
)

func TestValidEmail(t *testing.T) {
	want := "qwikle@gmail.com"

	got, err := value_objects.ValidEmail(want)
	if err != nil {
		t.Fatalf("ValidEmail(%q) = %q, want %q", want, got, want)
	}
}

func TestInvalidEmail(t *testing.T) {
	want := "qwikle@gmail"

	_, err := value_objects.ValidEmail(want)
	if err == nil {
		t.Fatalf("ValidEmail(%q) = %q, want %q", want, err, want)
	}
}

func TestEmptyEmail(t *testing.T) {
	want := ""

	_, err := value_objects.ValidEmail(want)
	if err == nil {
		t.Fatalf("ValidEmail(%q) = %q, want %q", want, err, want)
	}
}
