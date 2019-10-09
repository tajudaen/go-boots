package main

import (
	"fmt"
)

func main() {
	s := "this is a string"
	fmt.Printf("%v, %T\n", s, s)

	// rune (alias for int32)
	r := 'a'
	fmt.Printf("%v, %T\n", r, r)
}