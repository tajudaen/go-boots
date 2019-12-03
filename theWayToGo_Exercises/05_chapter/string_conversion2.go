package main
import (
	"fmt"
	"strconv"
)

func main()  {
	var orig string = "1234"

	c, _ := Converter(orig)

	fmt.Printf("The integer is %d\n", c)
}

func Converter(s string) (int, error) {
	an, err := strconv.Atoi(s)

	if err != nil {
		fmt.Printf("s %s is not an integer - exiting with error\n", s)
		return 0, err
	}

	return an, nil
}