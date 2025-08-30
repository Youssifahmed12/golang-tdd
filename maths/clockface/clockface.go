package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

const secondHandLength = 90
const clockCentreX = 150
const clockCentreY = 150

func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	return Point{p.X*secondHandLength + clockCentreX, -p.Y*secondHandLength + clockCentreY}
}
func secondsInRadians(t time.Time) float64 {
	return (float64(t.Second()) / 30) * math.Pi
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
