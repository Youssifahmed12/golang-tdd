package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("Adding Two Positive Numbers", func(t *testing.T) {
		sum := Add(6, 8)
		expected := 14
		if sum != expected {
			t.Errorf("Wanted %d got %d", expected, sum)
		}
	})
}

func ExampleAdd() {
	sum := Add(5, 5)
	fmt.Println(sum)
	// Output: 10
}
