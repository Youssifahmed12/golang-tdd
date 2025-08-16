package main

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 6, 7, 5, 3}
	sum := Sum(numbers)
	expected := 22
	if sum != expected {
		t.Errorf("got %d wanted %d given %v", sum, expected, numbers)
	}
}
