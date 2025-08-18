package main

import (
	"bytes"
	"slices"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 then go", func(t *testing.T) {

		bytes := &bytes.Buffer{}
		ss := &SpyCountDownOperation{}
		Countdown(bytes, ss)

		got := bytes.String()
		want := `3
2
1
Go!
`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})
	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountDownOperation{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !slices.Equal(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}

	})
}

func TestConfigurableSleeper(t *testing.T) {
	spyTime := SpyTime{}
	duartionSleep := 5 * time.Second
	confSleeper := ConfigurableSleeper{duartionSleep, spyTime.SetDurationSlept}
	confSleeper.Sleep()

	got := confSleeper.duration
	want := spyTime.durationSlept

	if got != want {
		t.Errorf("got %v , want %v ", got, want)
	}
}

const (
	write = "write"
	sleep = "sleep"
)

type SpyCountDownOperation struct {
	Calls []string
}

func (s *SpyCountDownOperation) Sleep() {
	s.Calls = append(s.Calls, sleep)
}
func (s *SpyCountDownOperation) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) SetDurationSlept(duration time.Duration) {
	s.durationSlept = duration
}
