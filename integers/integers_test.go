package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("Adding Two Positive Numbers", func(t *testing.T) {
		sum := add(6, 8)
		expected := 14
		if sum != expected {
			t.Errorf("Wanted %d got %d", expected, sum)
		}
	})
}

func Exampleadd() {
	sum := add(5, 5)
	fmt.Println(sum)
	// Output: 10
}
