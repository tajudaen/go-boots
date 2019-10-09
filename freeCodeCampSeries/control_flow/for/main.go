package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	a := 0
	for {
		fmt.Println(a)
		a++
		if a == 5 {
			break
		}
	}

	// Label
	Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i * j >= 3 {
				break Loop
			}
		}
	}

	s := []int{1,2,3}
	for _, v := range s {
		fmt.Println(v)
	}
}