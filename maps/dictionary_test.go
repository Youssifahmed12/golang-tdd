package maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run("find string available in dictionary", func(t *testing.T) {
		dic := Dictionary{"test": "this is a test"}

		want := "this is a test"
		got, _ := dic.Search("test")

		assertStrings(t, got, want)
	})
	t.Run("find non existent string in dictionary", func(t *testing.T) {
		dic := Dictionary{"nottest": "this is a test"}

		want := ErrWordDoesntExist
		_, got := dic.Search("test")
		assertError(t, got, want)

	})
}

func TestAdd(t *testing.T) {
	t.Run("adding to the dictionary", func(t *testing.T) {
		var (
			word          = "word"
			wordDefintion = "word definition"
		)

		dic := Dictionary{}
		dic.Add(word, wordDefintion)
		assertDefinition(t, dic, word, wordDefintion)
	})
}

func assertDefinition(t testing.TB, d Dictionary, word string, wordDef string) {
	t.Helper()
	got, err := d.Search(word)

	if err != nil {
		t.Fatal("should've found word")
	}
	assertStrings(t, got, wordDef)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
