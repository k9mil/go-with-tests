package main

import "testing"

// func TestHello(t *testing.T) {
// 	t.Run("Saying hello to people", func(t *testing.T) {
// 		got := Hello("Joe")
// 		want := "Hello, Joe"

// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	})
// 	t.Run("Saying 'Hello, World' when aj empty string is supplied", func(t *testing.T) {
// 		got := Hello("")
// 		want := "Hello, World"

// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}
// 	})
// }

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("Topuria", "Spanish")
		want := "Hola, Topuria"
		assertCorrectMessage(t, got, want)
	})
}
