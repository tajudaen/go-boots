package main

import (
	"fmt"
	"time"
	"runtime"
)

// Concurrency is handling or dealing with multiple things or processes at the same time i.e doing something and while waiting for response of a process, instead of been idle, we move to the other process and go back to the previous process when there is at an available time

// Parallelism is doing multiple things all at once or at a go


func main() {
	godur, _ := time.ParseDuration("10ms")
	runtime.GOMAXPROCS(2) // GOMAXPROCS adds Parallelism to the application. GOMAXPROCS(2) allows the application to use up to 2 processors(CPU) for the application

	// Add concurrency by adding the word (go) to the front of the function
	go func ()  {
		for i := 0; i < 100; i++ {
			fmt.Println("Hello")
			time.Sleep(godur)	// This puts the current goroutine to sleep temporary, so as to allow the application work on the other goroutine
		}
	}()

	go func ()  {
		for i := 0; i < 100; i++ {
			fmt.Println("Go")
			time.Sleep(godur) // This puts the current goroutine to sleep temporary, so as to allow the application work on the other goroutine
		}
	}()

	// Put the program to sleep for 1second, so that the two other goroutines can be called before the application(main routine) exit
	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}