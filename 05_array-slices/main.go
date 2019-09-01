package main
import "fmt"

func main() {
	// ARRAYS - fixed length
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Declare and Assign
	namesArr := [2]string{"John", "Doe"}

	// SLICES - no fixed length
	stack := []string{"PHP", "NodeJs", "Golang"}

	fmt.Println(fruitArr)
	fmt.Println(namesArr)
	fmt.Println(stack)
}