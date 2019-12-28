package arrays

import "testing"

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := [6]int{1, 1, 1, 1, 1,1}
		got := Sum(numbers)
		want := 6
		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}
