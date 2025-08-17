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

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 10.0}, 100.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g wanted %g", got, tt.want)
		}
	}

}
