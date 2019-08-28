package main

import "fmt"

func main() {
	// Define map (a map is a key value pair)
	emails := make(map[string]string)

	// Assign
	emails["john"] = "john@doe.com"
	emails["red"] = "fox@red.com"
	emails["if"] = "maps@if.com"

	// Declare and Assign
	integers := map[int]string{1: "one", 2: "two", 3: "three"}

	fmt.Println(emails)
	fmt.Println(integers)

	// Delete from map
	delete(emails, "if")

	fmt.Println(emails)
}