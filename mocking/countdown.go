package mocking

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SypSleeper struct {
	Calls int
}

func (s *SypSleeper) Sleep() {
	s.Calls++
}

type ConfigurableSleeper struct {
	duration time.Duration
}

func (c *ConfigurableSleeper) Sleep() {
	time.Sleep(c.duration)
}

const finalWord = "Go!"
const CountdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := CountdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}
