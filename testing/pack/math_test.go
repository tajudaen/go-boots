package pack

import (
	"testing"
	"fmt"
)

var addTable = []struct {
	in  []int
	out int
}{
	{[]int{1, 2}, 3},
	{[]int{1, 2, 3, 4}, 10},
}

func TestCanAddNumbers(t *testing.T) {
	for _, entry := range addTable {
		result := Add(entry.in...)
		if result != entry.out {
			t.Error("Failed to add numbers as expected")
		}
	}

	result := Add(1, 2)

	if result != 3 {
		t.Log("Failed to add one and two")
		t.Fail()
	}
}

func TestCanSubtractNumbers(t *testing.T) {
	result := Subtract(10, 5)

	if result != 5 {
		t.Error("Failed to subtract two numbers properly")
	}
}

func TestCanMultiplyNumbers(t *testing.T) {
	t.Skip("Not implemented")
}

func Example() {
	result := Add(1, 2)

	fmt.Println(result)

	// Output:
	// 31
}