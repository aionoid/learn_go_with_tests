// @Title
// @Description
// @Author
// @Update

package main

import (
	"bytes"
	"reflect"
	"testing"
)

const sleep = "sleep"
const write = "write"

type SpyCountDownOpperation struct {
	Calls []string
}

func (s *SpyCountDownOpperation) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountDownOpperation) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("print 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleep := &SpyCountDownOpperation{}

		Countdown(buffer, spySleep)
		got := buffer.String()
		want := "3\n2\n1\nGo!"

		if got != want {
			t.Errorf("got %q and want %q", got, want)
		}
		// if spySleep.Calls != 3 {
		//   t.Errorf("not enough calls to sleeper, want 3 and got %d", spySleep.Calls)
		// }
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountDownOpperation{}

		Countdown(spySleepPrinter, spySleepPrinter)
		got := spySleepPrinter.Calls
		want := []string{write, sleep, write, sleep, write, sleep, write}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got calls %v but wanted %v", got, want)
		}
	})
}
