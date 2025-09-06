package iterations

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 8)
	expected := "aaaaaaaa"
	if repeated != expected {
		t.Errorf("got %q , should've got %q", repeated, expected)
	}
}
func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 8)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 6)
	fmt.Println(repeated)
	// Output: aaaaaa
}
