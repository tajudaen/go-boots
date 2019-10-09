package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	var msg = "Hello"
	wg.Add(1)
	go say(msg)
	msg = "Goodbye"
	wg.Wait()
}

func say(msg string) {
	fmt.Println(msg)
	wg.Done()
}