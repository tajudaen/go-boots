package main

import "fmt"

func main() {
	n := 42
	fmt.Printf("%v, %T\n", n, n)

	// bit operators (only available on integer numbers)
	a := 10 // 1010
	b := 3 // 0011

	fmt.Println(a & b) // 0010 = 2 (when both position is set i.e 1)
	fmt.Println(a | b) // 1011 = 11 (one of them is set at that position)
	fmt.Println(a ^ b) // 1001 = 9 (one of them have a bit/position set but not both)
	fmt.Println(a &^ b) // 0100 = 8 (opposite of | operator)

	// bit shifting (only available on integer numbers)
	c := 8 // 2^3
	fmt.Println(c << 3) // 2^3 * 2^3 = 2^6 = 64
	fmt.Println(c >> 3) // 2^3 / 2^3 = 2^0 = 1

	// complex numbers(a number with both real and imaginary parts)
	d := 1 + 2i // 1 - real part, 2i - imaginary part
	e := 2 + 5.2i
	
	fmt.Println(d + e)
	fmt.Println(d - e)
	fmt.Println(d * e)
	fmt.Println(d / e)

	var m complex64 = complex(5, 12)
	fmt.Printf("%T, %v\n", m, m)
	fmt.Printf("%T, %v\n", real(m), real(m))
	fmt.Printf("%T, %v\n", imag(m), imag(m))
}
