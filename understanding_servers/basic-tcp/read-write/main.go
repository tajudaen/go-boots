package main

import (
	"fmt"
	"bufio"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}
func handle(conn net.Conn)  {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintln(conn, "I heard you say", ln)
	}
		defer conn.Close() 

		// we never get here
		// we have an open stream connetion
		// how does the above reader know when it's done
		fmt.Fprintln(conn, "Code got here.")
}