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
	checkArea := func(t testing.TB, s Shape, want float64) {
		t.Helper()
		got := s.Area()
		if got != want {
			t.Errorf("got %.2f wanted %.2f", got, want)
		}
	}
	t.Run("Calculate the Area of a rectangle", func(t *testing.T) {
		r := Rectangle{10.0, 10.0}
		want := 100.0
		checkArea(t, r, want)
	})
	t.Run("Calculate the Area of a circle", func(t *testing.T) {
		c := Circle{10}
		want := 314.1592653589793
		checkArea(t, c, want)
	})

}
