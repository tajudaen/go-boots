package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("This is true")
	}

	if b := false; !b {
		fmt.Println(b)
	}
}