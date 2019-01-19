package maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run("If the word is a match", func(t *testing.T) {
		dictionary := map[string]string{"test": "this is just a test"}

		got := Search(dictionary, "test")
		want := "this is just a test"

		if got != want {
			t.Errorf("Expected: %s, Wanted: %s", want, got)
		}
	})

	t.Run("If Word wasn't found", func(t *testing.T) {
		dictionary := map[string]string{"test": "this is just a test"}

		got := Search(dictionary, "math")
		want := "Couldn't find the word"

		if got != want {
			t.Errorf("Expected: %s, Wanted: %s", want, got)
		}
	})

}
