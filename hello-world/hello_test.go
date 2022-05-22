package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func (t *testing.T) {
		got := Hello("Mark", "Spanish")
		want := "Hola, Mark"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func (t *testing.T) {
		got := Hello("Mark", "French")
		want := "Bonjour, Mark"
		assertCorrectMessage(t, got, want)
	})
}