package main

import (
	"fmt"
)

type myStruct struct {
	foo int
}

func main()  {
	var a int = 42
	var b *int = &a
	fmt.Println(&a, b)
	*b = 14
	fmt.Println(a, *b)

	c := [3]int{1, 2, 3}
	d := &c[0]
	e := &c[1]
	fmt.Printf("%v %p %p\n", c, d, e)

	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms)
}