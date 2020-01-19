package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Word counter has started, input your word")
	input, err := inputReader.ReadString('S')

	if err != nil {
		fmt.Println(err)
	}

	
	line := strings.Count(input, string('\n'))

	fmt.Println("Number of lines", line)

	c := strings.ReplaceAll(input, string('\n'), "")
	fmt.Println(len(c))

}