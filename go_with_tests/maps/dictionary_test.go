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
		err := dictionary.Add("new", "a new addition")
		if err != nil {
			t.Error("failed to add new word to dictionary")
		}
		expected := "a new addition"
		assertDefinition(t, dictionary, "new", expected)
	})

	t.Run("Add an already existing word to a dictionary", func(t *testing.T) {
		got := dictionary.Add("new", "the same new addition")
		expected := ErrTermExists
		assertError(t, got, expected)
	})
}

func TestUpdate(t *testing.T) {
	dictionary := Dictionary{}

	t.Run("Update an existing term", func(t *testing.T) {
		err := dictionary.Add("test", "original definition")
		if err != nil {
			t.Error("failed to add new word to dictionary")
		}
		err = dictionary.Update("test", "new definition")
		if err != nil {
			t.Error("failed to update definition")
		}
		assertDefinition(t, dictionary, "test", "new definition")
	})

	t.Run("Update a non-existent term", func(t *testing.T) {
		err := dictionary.Update("new test", "this should not work")
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{}

	t.Run("Delete an existing term", func(t *testing.T) {
		err := dictionary.Add("test", "original definition")
		if err != nil {
			t.Error("failed to add new word to dictionary")
		}
		dictionary.Delete("test")
		_, err = dictionary.Search("test")
		if err == nil {
			t.Error("term not deleted from dictionary")
		}
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
	got, err := dict.Search(term)
	if err != nil {
		t.Fatal("should find the added term", err)
	}

	assertStrings(t, got, definition)
}
