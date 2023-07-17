package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("world")  // Invoke our function
	want := "Hello, world" // hard code expected value

	if got != want {
		// We are calling the Errorf method on our t which will print out a message and fail the test.
		// The f stands for format which allows us to build a string with values inserted into the placeholder values %q
		t.Errorf("got %q want %q", got, want)
	}
}

func TestSpecificHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
