package arrays

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 1, 1, 1, 1}
	got := Sum(numbers)
	want := 5
	if want != got {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
