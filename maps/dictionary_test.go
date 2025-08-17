package maps

import "testing"

func TestMap(t *testing.T) {
	t.Run("find string available in dictionary", func(t *testing.T) {
		dic := Dictionary{"test": "this is a test"}

		want := "this is a test"
		got, _ := dic.Search(dic, "test")

		assertStrings(t, got, want)
	})
	t.Run("find non existent string in dictionary", func(t *testing.T) {
		dic := Dictionary{"nottest": "this is a test"}

		want := ErrWordDoesntExist
		_, got := dic.Search(dic, "test")
		assertError(t, got, want)

	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
