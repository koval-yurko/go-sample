package sample

import (
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hi, aa. Welcome!"
	if got := Hello("aa"); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
