package mocking

import (
	"fmt"
	"io"
)

const finalWord = "Go!"
const CountdownStart = 3

func Countdown(out io.Writer) {
	for i := CountdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}

	fmt.Fprint(out, finalWord)
}
