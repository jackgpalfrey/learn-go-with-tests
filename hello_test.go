package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to specfic name", func(t *testing.T) {
		name := "Jack"
		got := Hello(name)
		want := "Hello, " + name
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to world", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)

	})
}


func assertCorrectMessage(t testing.TB, got string, want string){
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
