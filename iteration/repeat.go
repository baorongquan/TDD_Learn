package iteration

import (
	"strings"
)

// Repeat returns a new string consisting of count copies of the string character
func Repeat(character string, count int) string {
	var repeated string
	repeated = strings.Repeat(character, count)
	return repeated
}
