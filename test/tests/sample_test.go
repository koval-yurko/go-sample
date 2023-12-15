package tests

import (
	"testing"

	"github.com/koval-yurko/go-sample/client/sample"
)

func TestHello(t *testing.T) {
	want := "Hi, aa. Welcome v0.1.0!"
	if got := sample.Hello("aa"); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
