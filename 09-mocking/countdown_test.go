// @Title
// @Description
// @Author
// @Update

package main

import (
	"bytes"
	"testing"
)

type SpySleep struct {
	Calls int
}

func (s *SpySleep) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleep := &SpySleep{}

	Countdown(buffer, spySleep)
	got := buffer.String()
	want := "3\n2\n1\nGo!"

	if got != want {
		t.Errorf("got %q and want %q", got, want)
	}

	if spySleep.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3 and got %d", spySleep.Calls)
	}
}
