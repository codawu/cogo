package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 1, 1, 1, 1, 1}
		got := Sum(numbers)
		want := 6
		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	// slice can not compare use ==
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumTails(t *testing.T) {
	got := SumTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
