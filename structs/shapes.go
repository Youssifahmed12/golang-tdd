package structs

type Rectangle struct {
	height float64
	width  float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.height + r.width)
}

func Area(r Rectangle) float64 {
	return r.height * r.width
}
