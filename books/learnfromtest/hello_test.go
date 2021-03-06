package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		const name = "Bob"
		got := Hello("Bob", "")
		want := "Hello, " + name
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		const name = "Bob"
		got := Hello("Bob", "Spanish")
		want := "Hola, " + name
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		const name = "Bob"
		got := Hello("Bob", "French")
		want := "Bonjour, " + name
		assertCorrectMessage(t, got, want)
	})

}
