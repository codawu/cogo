package structs

import (
	"reflect"
	"testing"
)

func TestShape(t *testing.T) {

	t.Run("shape perimeter", func(t *testing.T) {
		shapes := []Shape{
			Rectangle{4, 5},
			Triangle{3, 4, 5},
			Circle{5},
		}
		want := []float64{(4 + 5) * 2, 12, 2 * 5 * PI}
		var got []float64
		for _, shape := range shapes {
			got = append(got, shape.Perimeter())
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v but got %v", want, got)
		}
	})

	t.Run("shape area", func(t *testing.T) {
		shapes := []Shape{
			Rectangle{4, 5},
			Triangle{3, 4, 5},
			Circle{5},
		}
		want := []float64{4 * 5, 6, 5 * 5 * PI}
		var got []float64
		for _, shape := range shapes {
			got = append(got, shape.Area())
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v but got %v", want, got)
		}
	})

}
