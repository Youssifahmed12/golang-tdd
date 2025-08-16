package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("English Hello Test with name", func(t *testing.T) {
		want := "Hello Youssif!"
		got := hello("Youssif", "English")
		if want != got {
			t.Errorf("wanted %q but got %q", want, got)
		}
	})
	t.Run("Spanish Hello Test with name", func(t *testing.T) {
		want := "Hola Youssif!"
		got := hello("Youssif", "Spanish")
		if want != got {
			t.Errorf("wanted %q but got %q", want, got)
		}
	})
	t.Run("French Hello Test with name", func(t *testing.T) {
		want := "Bonjour Youssif!"
		got := hello("Youssif", "French")
		if want != got {
			t.Errorf("wanted %q but got %q", want, got)
		}
	})
}
