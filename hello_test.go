package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to specfic name", func(t *testing.T) {
		name := "Jack"
		got := Hello(name)
		want := "Hello, " + name

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("saying hello to world", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
