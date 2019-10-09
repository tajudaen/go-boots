package main

import (
	"fmt"
)

func main() {
	var i int
	i = 42
	j := 7.
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T", j, j)
}