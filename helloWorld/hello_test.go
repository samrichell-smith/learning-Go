package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Sam", "")
		want := "Hello, Sam"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Espanol", func(t *testing.T) {
		got := Hello("Juan", "Espanol")
		want := "Hola, Juan"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Ameilie", "French")
		want := "Bonjour, Ameilie"
		assertCorrectMessage(t, got, want)
	})

	
}

func assertCorrectMessage(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
}
