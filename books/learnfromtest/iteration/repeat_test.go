package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeat := Repeat("hello",5)
	except := "hellohellohellohellohello"

	if except != repeat {
		t.Errorf("excepted '%q' but got '%q'", except, repeat)
	}
}
