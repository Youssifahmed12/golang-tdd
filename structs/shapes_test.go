package structs

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("Calculate the perimeter of a rectangle", func(t *testing.T) {
		r := Rectangle{10.0, 10.0}
		got := Perimeter(r)
		want := 40.0
		if got != want {
			t.Errorf("got %.2f wanted %.2f", got, want)
		}
	})
}
func TestArea(t *testing.T) {
	t.Run("Calculate the Area of a rectangle", func(t *testing.T) {
		r := Rectangle{10.0, 10.0}
		got := Area(r)
		want := 100.0
		if got != want {
			t.Errorf("got %.2f wanted %.2f", got, want)
		}
	})
}
