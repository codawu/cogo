package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		const name = "Bob"
		got := Hello(name)
		want := "Hello, " + name
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})

	t.Run("saying hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})

}
