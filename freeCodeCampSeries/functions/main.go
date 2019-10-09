package main

import (
	"fmt"
)

type greeter struct {
	greeting string
	name string
}

func main() {
	sayMessage("Hello Go")
	sum(1, 2, 3, 4, 5)
	d, err := divide(5.0,2.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
	// Anonymous function
	func ()  {
		fmt.Println("I'm anonymous")
	}()
}

func sayMessage(msg string)  {
	fmt.Println(msg)
} 

func sum(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}