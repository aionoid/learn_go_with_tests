package hello

// @Title
// @Description
// @Author
// @Update
import "testing"

func TestHello(t *testing.T) {
	t.Run("welcome message for someone", func(t *testing.T) {
		want := "Hello, Jack!"
		got := Hello("Jack", "eng")

		assertMessage(want, got, t)
	})

	t.Run("Default Welcome if no name is supplied", func(t *testing.T) {
		want := "Hello, World!"
		got := Hello("", "eng")

		assertMessage(want, got, t)
	})
	t.Run("welcome message in frech", func(t *testing.T) {
		want := "Bonjour, Jack!"
		got := Hello("Jack", "fr")

		assertMessage(want, got, t)
	})
	t.Run("welcome message in spanish", func(t *testing.T) {
		want := "Holla, Jack!"
		got := Hello("Jack", "sp")

		assertMessage(want, got, t)
	})
}

func assertMessage(want, got string, t testing.TB) {
	t.Helper()
	if want != got {
		t.Errorf("wanted %q, but got %q", want, got)
	}
}
