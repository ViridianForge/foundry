package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris!"

		assertCorrectHello(t, got, want)
	})
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectHello(t, got, want)
	})
	t.Run("saying hello in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie!"

		assertCorrectHello(t, got, want)
	})
	t.Run("saying hello in French", func(t *testing.T) {
		got := Hello("Benoit", "French")
		want := "Bonjour, Benoit!"

		assertCorrectHello(t, got, want)
	})
}

func assertCorrectHello(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, expected %q", got, want)
	}

}
