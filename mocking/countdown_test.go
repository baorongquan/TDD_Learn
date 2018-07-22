package mocking

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SypSleeper{}
	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not enough call to sleeper, want 4 got %d", spySleeper.Calls)
	}
}
