package main

import (
	"fmt"
	"strings"
)

func main() {
	phrase := "These are the times that try men's souls.\n"

	words := strings.Split(phrase, " ")

	ch := make(chan string, len(words)) // buffered channel; the second argument to make() is the storage capacity of the channel

	for _, word := range words {
		ch <- word
	}

	close(ch) // closing a channel; a closed channel cannot receive new message once its closed

	for i := 0; i < len(words); i++ {
		fmt.Print(<-ch + " ")
	}
}