package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func handleConnection(c net.Conn) {

	fmt.Printf("Serveing %s\n", c.RemoteAddr().String())

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')

		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed")
				break
			}
			log.Fatal(err)
			return
		}

		temp := strings.TrimSpace(netData)

		c.Write([]byte(temp + "\n"))
	}

	// Close connection when this function ends
	defer c.Close()
}

func main() {
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)
	}
}
