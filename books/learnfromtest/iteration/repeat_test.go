package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeat := Repeat("hello")
	except := "hellohello"

	if except != repeat {
		t.Errorf("excepted '%q' but got '%q'", except, repeat)
	}
}
