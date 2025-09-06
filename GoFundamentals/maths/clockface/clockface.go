package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

// used int here bec the Minute / Second function returns an int
func timeToRadain(t int) float64 {
	return (float64(t) / 30) * math.Pi
}

func secondsInRadians(t time.Time) float64 {
	return timeToRadain(t.Second())
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / 60) +
		timeToRadain(t.Minute())
}

func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / 12) +
		(math.Pi / (6 / float64(t.Hour()%12)))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}
