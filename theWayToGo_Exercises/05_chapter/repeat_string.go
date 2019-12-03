package main

import (
	"strings"
)

// Returns a string in a pyramid structure
func Repeat(s string, n int) string {
	var result string
	for i := 1; i <= n; i++ {
		result += strings.Repeat(s, i)
		result += "\n"
	}
	return result
}
