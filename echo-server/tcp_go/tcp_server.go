package main

import (
	"fmt"
	"io"
	"log"
	"net"
	_ "net/http/pprof"
	"syscall"
)

func main() {
	checkKernelParam()

	ln, err := net.Listen("tcp", ":8972")
	if err != nil {
		panic(err)
	}

	var connections []net.Conn

	defer func() {
		for _, conn := range connections {
			conn.Close()
		}
	}()

	for {
		conn, e := ln.Accept()
		if e != nil {
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				log.Printf("accept temp err: %v", ne)
				continue
			}

			log.Printf("accept err: %v", e)
			return
		}

		go handleConn(conn)

		connections = append(connections, conn)

		if len(connections)%100 == 0 {
			log.Printf("total number of connections: %v", len(connections))
		}
	}
}

func handleConn(conn net.Conn) {
	io.Copy(io.Discard, conn)
}

func checkKernelParam() {
	// see SOMMAXCONN
	fmt.Println("Currnet SOMAXCONN: ", syscall.SOMAXCONN)
}
