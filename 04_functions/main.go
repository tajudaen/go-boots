package main

import (
	"fmt"
	"strings"
)

func greeting(name string) string {
	return "Hello " + name
}

func main() {
	fmt.Println(greeting("Tajudeen"))

	module := "functions"
	author := "john doe"

	fmt.Println(converter(module, author))

	bestFinish := bestLeagueFinishes(13, 10, 13, 17, 14, 16)
	fmt.Println(bestFinish)
}

func converter(s1, s2 string) (str1, str2 string) {
	s1 = strings.Title(s1)
	s2 = strings.ToUpper(s2)

	return s1, s2
}

//Variadic function
func bestLeagueFinishes(finishes ...int) int {
	best := finishes[0]

	for _, position := range finishes {
		if position < best {
			best = position
		}
	}
	return best
}