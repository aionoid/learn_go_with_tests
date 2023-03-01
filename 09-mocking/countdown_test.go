// @Title
// @Description
// @Author
// @Update

package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const sleep = "sleep"
const write = "write"

type SpyCountDownOpperation struct {
	Calls []string
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
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

func TestConfigurableSleeper(t *testing.T) {
	sleeptime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleeptime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleeptime {
		t.Errorf("should have slept for %v but slept for %v", sleeptime, spyTime.durationSlept)
	}
}
