package iterations

import "testing"

func TestRepeat(t *testing.T) {
	repeated := repeat("a")
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("got %q , should've got %q", repeated, expected)
	}
}
func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		repeat("a")
	}
}
