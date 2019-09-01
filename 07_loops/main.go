package main

import "fmt"

func main() {
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	for i := 5; i < 10; i++ {
		fmt.Printf("Number &d\n", i)		
	}

	// FizzBuzz
	for i := 1; i < 30; i++ {
		if i % 15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}