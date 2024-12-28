package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	for i := 0; i < 100_000; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")

		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(conn, "Hello, Server\n")

		status, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Message from server: %s", status)

		// Close connection when this function ends
		defer conn.Close()
	}
}
