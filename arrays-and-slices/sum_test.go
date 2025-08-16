package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum slice of size 5 ", func(t *testing.T) {
		numbers := []int{1, 6, 7, 5, 3}
		sum := Sum(numbers)
		expected := 22
		if sum != expected {
			t.Errorf("got %d wanted %d given %v", sum, expected, numbers)
		}
	})
}

// reflect.DeepEqual doesnt type check but works with every type ,
// on the other hand slices has a function to compare only slices
func TestSumAll(t *testing.T) {
	t.Run("sum all slices using deep equal", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{1, 6, 3})
		want := []int{6, 10}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v wanted %v", got, want)
		}
	})
	t.Run("sum all slices using deep equal", func(t *testing.T) {

		got := SumAll([]int{1, 2, 3}, []int{1, 6, 3})
		want := []int{6, 10}
		if !slices.Equal(got, want) {
			t.Errorf("got %v wanted %v", got, want)
		}
	})
}
