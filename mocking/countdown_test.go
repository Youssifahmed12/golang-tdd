package main

import (
	"bytes"
	"slices"
	"testing"
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
