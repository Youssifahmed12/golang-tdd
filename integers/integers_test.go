package integers

import "testing"

func TestAdd(t *testing.T) {
	t.Run("Adding Two Positive Numbers", func(t *testing.T) {
		var (
			num1 = 5
			num2 = 9
		)
		want := 14
		got := add(num1, num2)
		if want != got {
			t.Errorf("Wanted %b got %b", want, got)
		}
	})

}
