// @Title
// @Description
// @Author
// @Update
package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type DefaultSleeper struct{}

func (d DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type Sleeper interface {
	Sleep()
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprint(writer, finalWord)
}

func main() {
	fmt.Println("vim-go")
	slp := &DefaultSleeper{}
	Countdown(os.Stdout, slp)
}
