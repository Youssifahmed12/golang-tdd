package maps

import "testing"

func TestMap(t *testing.T) {
	dic := map[string]string{"test": "this is a test"}
	want := "this is a test"
	got := Search(dic, "test")

	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}
