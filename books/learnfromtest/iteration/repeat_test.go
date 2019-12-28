package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeat := Repeat("hello", 5)
	except := "hellohellohellohellohello"

	if except != repeat {
		t.Errorf("excepted '%q' but got '%q'", except, repeat)
	}
}

func BenchmarkRepeat(b *testing.B) {
	times := 5
	for i := 0; i < b.N; i++ {
		Repeat("Hello", times)
	}
}
