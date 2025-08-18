package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (cs *ConfigurableSleeper) Sleep() {
	cs.sleep(cs.duration)
}

const (
	countdownStart = 3
	finalWord      = "Go!"
)

func Countdown(writer io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		s.Sleep()
	}
	fmt.Fprintln(writer, finalWord)
}

func main() {
	cs := &ConfigurableSleeper{3 * time.Second, time.Sleep}
	Countdown(os.Stdout, cs)
}
