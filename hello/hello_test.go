package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("English Hello Test with name", func(t *testing.T) {
		want := "Hello Youssif!"
		got := hello("Youssif", "English")
		assertCorrectGreeting(t, want, got)
	})
	t.Run("Spanish Hello Test with name", func(t *testing.T) {
		want := "Hola Youssif!"
		got := hello("Youssif", "Spanish")
		assertCorrectGreeting(t, want, got)
	})
	t.Run("French Hello Test with name", func(t *testing.T) {
		want := "Bonjour Youssif!"
		got := hello("Youssif", "French")
		assertCorrectGreeting(t, want, got)
	})
	t.Run("Undefined Hello Test with name", func(t *testing.T) {
		want := "Hello Youssif!"
		got := hello("Youssif", "Zimbabwian")
		assertCorrectGreeting(t, want, got)
	})

}

func assertCorrectGreeting(t testing.TB, want, got string) {
	t.Helper()
	if want != got {
		t.Errorf("wanted %q but got %q", want, got)
	}
}
