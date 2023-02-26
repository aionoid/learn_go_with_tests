// @Title
// @Description
// @Author
// @Update
package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Jack")

	got := buffer.String()
	want := "Hello, Jack!"

	if got != want {
		t.Errorf("got %q and want %q", got, want)
	}
}
