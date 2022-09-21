package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is only a test"}

	got := dictionary.Search("test")
	expected := "this is only a test"

	assertStrings(t, got, expected)
}

func assertStrings(t testing.TB, got string, expected string) {
	t.Helper()

	if got != expected {
		t.Errorf("got %q, expected %q", got, expected)
	}
}
