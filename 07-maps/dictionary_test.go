// @Title Dictionary tests
// @Description learn maps for go
// @Author Belaoura Yacoub
// @Update v0.0.1
package maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dict := Dictionary{"test": "this is just a test"}
		got, err := dict.Search("test")
		want := "this is just a test"

		assertError(t, err, nil)
		assertString(t, got, want)

	})
	t.Run("unknown word", func(t *testing.T) {
		dict := Dictionary{"test": "this is just a test"}
		_, err := dict.Search("test2")
		assertError(t, err, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {

	t.Run("add new word", func(t *testing.T) {
		dict := Dictionary{}
		dict.Add("test", "this is a new word")

		got, _ := dict.Search("test")
		want := "this is a new word"
		assertString(t, got, want)
	})
	//
	t.Run("add an existing word", func(t *testing.T) {
		dict := Dictionary{"test": "this is a test"}
		err := dict.Add("test", "this is a new test")

		got, _ := dict.Search("test")
		want := "this is a test"

		assertError(t, err, ErrWordExiste)
		assertString(t, got, want)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("word existe", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"
		dictionary.Update(word, newDefinition)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("word dose not existe", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}
		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)
		assertError(t, err, ErrWordDoseNotExiste)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{word: definition}
	dictionary.Delete(word)
	_, err := dictionary.Search(word)
	assertError(t, err, ErrNotFound)
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, newDefinition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertString(t, got, newDefinition)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q but want %q", got, want)
	}
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q and want %q", got, want)
	}
}
