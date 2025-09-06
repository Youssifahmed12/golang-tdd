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

		want := ErrNotFound
		_, got := dic.Search("test")
		assertError(t, got, want)

	})
}

func TestAdd(t *testing.T) {
	t.Run("adding to the dictionary", func(t *testing.T) {

		dic := Dictionary{}
		word := "word"
		wordDefintion := "word definition"

		err := dic.Add(word, wordDefintion)
		assertError(t, err, nil)
		assertDefinition(t, dic, word, wordDefintion)
	})

	t.Run("adding a word that already exists", func(t *testing.T) {
		dic := Dictionary{"word": "doesn't matter"}
		word := "word"
		wordDefintion := "word definition"
		err := dic.Add(word, wordDefintion)
		if err == nil {
			t.Fatal("expected an error didnt get one")
		}

	})
}

func TestUpdate(t *testing.T) {
	t.Run("update existing word", func(t *testing.T) {
		dic := Dictionary{"word1": "word1 old description"}
		word := "word1"
		newDefinition := "word11 new description"

		err := dic.Update(word, newDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dic, word, newDefinition)

	})
	t.Run("update a word that doesn't exist", func(t *testing.T) {
		dic := Dictionary{}
		word := "word"
		newDefinition := "a simple description"

		err := dic.Update(word, newDefinition)
		assertError(t, err, ErrNotFound)
	})

}

func TestDelete(t *testing.T) {
	t.Run("Delete existing item", func(t *testing.T) {
		dic := Dictionary{"word": "word definition"}

		err := dic.Delete("word")

		assertError(t, err, nil)

		_, err = dic.Search("word")

		assertError(t, err, ErrNotFound)

	})

	t.Run("Delete non existing item", func(t *testing.T) {
		dic := Dictionary{}
		err := dic.Delete("word")
		assertError(t, err, ErrCouldntDelete)
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
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
