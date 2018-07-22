package dependencyInjection

import (
	"fmt"
	"io"
)

func Greet(write io.Writer, name string) {
	fmt.Fprintf(write, "Hello, %s", name)
}
