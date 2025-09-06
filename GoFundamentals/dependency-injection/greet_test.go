package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	bytes := bytes.Buffer{}
	Greet(&bytes, "Youssif was here")

	got := bytes.String()
	want := "Hello, Youssif was here"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
