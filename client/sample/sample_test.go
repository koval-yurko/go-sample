package sample

import (
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hi, aa. Welcome v0.1.0!"
	if got := Hello("aa"); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
