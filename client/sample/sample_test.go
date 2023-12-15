package sample

import (
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hi, aa. Welcome v0.0.6!"
	if got := Hello("aa"); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
