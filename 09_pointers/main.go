package main

import "fmt"

func main() {
	a := 5

	// & is a pointer to a memory address; assigning (b) to a pointer of (a)
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// Use * to read value from address; (*b) accessing the value in that memory location
	fmt.Println(*b)

	// Change value with pointer
	*b = 10
	fmt.Println(a)
}