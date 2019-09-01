package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}

// METHODS(value receiver and pointer receiver)

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver - used to change values in a struct)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p. lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "John", lastName: "Doe", city: "Lagos", gender: "m", age: 20}
	person2 := Person{"Jane", "Holmes", "Lagos", "f", 19}

	person1.hasBirthday()
	fmt.Println(person1)

	person2.getMarried("Doe")
	fmt.Println(person2.greet())
}