package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 1) // create a channel
	ch <- "Hello" // pass message into a channel

	fmt.Println(<-ch) // drain the channel
}