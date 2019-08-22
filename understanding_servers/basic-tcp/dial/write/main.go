package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println(conn, "I dialed you.")
}