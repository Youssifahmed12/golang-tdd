package maps

import "testing"

func TestMap(t *testing.T) {
	dic := Dictionary{"test": "this is a test"}
	want := "this is a test"
	got := dic.Search(dic, "test")
	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given", got, want)
	}
}
