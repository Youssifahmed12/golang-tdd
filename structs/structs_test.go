package structs

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("Calculate the perimeter of a rectangle", func(t *testing.T) {
		got := Perimeter(10.0, 10.0)
		want := 40.0
		if got != want {
			t.Errorf("got %.2f wanted %.2f", got, want)
		}
	})
}
