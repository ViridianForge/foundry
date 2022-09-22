package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is only a test"}

	t.Run("Valid Search", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		expected := "this is only a test"

		assertStrings(t, got, expected)
	})

	t.Run("Invalid Search", func(t *testing.T) {
		_, got := dictionary.Search("not_a_term")
		expected := ErrNotFound

		assertError(t, got, expected)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	t.Run("Add to an empty dictionary", func(t *testing.T) {
		dictionary.Add("new", "a new addition")
		expected := "a new addition"
		assertDefinition(t, dictionary, "new", expected)
	})
}

func assertStrings(t testing.TB, got string, expected string) {
	t.Helper()

	if got != expected {
		t.Errorf("got %q, expected %q", got, expected)
	}
}

func assertError(t testing.TB, got error, expected error) {
	t.Helper()

	if got != expected {
		t.Errorf("got %q, expected %q", got, expected)
	}
}

func assertDefinition(t testing.TB, dict Dictionary, term string, definition string) {
	t.Helper()
	got, err := dict.Search("new")
	if err != nil {
		t.Fatal("should find the added term", err)
	}

	assertStrings(t, got, definition)
}
