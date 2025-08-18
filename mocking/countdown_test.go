package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	bytes := &bytes.Buffer{}
	ss := &SpySleeper{}
	Countdown(bytes, ss)

	got := bytes.String()
	want := `3
2
1
Go!
`
	if ss.Calls != 3 {
		t.Errorf("sleeper called %d times instead of 3 times", ss.Calls)
	}
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

type SpySleeper struct {
	Calls int
}

func (ss *SpySleeper) Sleep() {
	ss.Calls++
}
