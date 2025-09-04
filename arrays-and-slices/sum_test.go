package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestReduce(t *testing.T) {
	t.Run("multiplying using reduce", func(t *testing.T) {
		multiply := func(result, number int) int {
			return result * number
		}

		AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concatenate", func(t *testing.T) {
		conc := func(result, number string) string {
			return result + number
		}

		AssertEqual(t, Reduce([]string{"a", "b", "c"}, conc, ""), "abc")
	})
}

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
	t.Run("sum all slices using slices equal", func(t *testing.T) {

		got := SumAll([]int{1, 2, 3}, []int{1, 6, 3})
		want := []int{6, 10}
		if !slices.Equal(got, want) {
			t.Errorf("got %v wanted %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSum := func(t testing.TB, sum, expected []int) {
		t.Helper()
		if !slices.Equal(sum, expected) {
			t.Errorf("got %v expected %v", sum, expected)
		}
	}
	t.Run("sum tails", func(t *testing.T) {
		sum := SumAllTails([]int{1, 2, 3, 4}, []int{2, 3, 0, 5, 6}, []int{9, 1, 0, 3})
		expected := []int{9, 14, 4}
		checkSum(t, sum, expected)
	})
	t.Run("safely sum empty tails", func(t *testing.T) {
		sum := SumAllTails([]int{}, []int{2, 3, 0, 5, 6}, []int{9, 1, 0, 3})
		expected := []int{0, 14, 4}
		checkSum(t, sum, expected)
	})

	t.Run("safelu sum one element slices", func(t *testing.T) {
		sum := SumAllTails([]int{}, []int{1}, []int{4})
		expected := []int{0, 1, 4}
		checkSum(t, sum, expected)
	})
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
