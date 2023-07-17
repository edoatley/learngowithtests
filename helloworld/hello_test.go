package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Pierre", "French")
		want := "Bonjour, Pierre"
		assertCorrectMessage(t, got, want)
	})

}

// In this helper function we've refactored our assertion into a new function.
// This reduces duplication and improves readability of our tests.
// We need to pass in t *testing.T so that we can tell the test code to fail when we need to.
//
// For helper functions, it's a good idea to accept a testing.TB which is an interface that *testing.T and *testing.B both satisfy,
// so you can call helper functions from a test, or a benchmark (don't worry if words like "interface" mean nothing to you right now, it will be covered later).
func assertCorrectMessage(t testing.TB, got, want string) {
	// t.Helper() is needed to tell the test suite that this method is a helper. By doing this when it fails the line number reported will be in our function call rather than inside our test helper.
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
