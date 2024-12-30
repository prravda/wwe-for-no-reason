package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"syscall"
	"time"
)

var (
	ip = flag.String("ip", "tcp_server", "server IP")
	//ip          = flag.String("ip", "localhost", "server IP")
	connections = flag.Int("conn", 10_000, "number of tcp connections")
)

func main() {
	flag.Parse()

	setLimit()

	addr := *ip + ":8972"
	log.Printf("Connect to %s", addr)

	var conns []net.Conn

	for i := 0; i < *connections; i++ {
		c, err := net.DialTimeout("tcp", addr, 10*time.Second)
		if err != nil {
			fmt.Println("failed to connect", i, err)
			i--
			continue
		}
		conns = append(conns, c)
		time.Sleep(time.Millisecond)
	}

	defer func() {
		for _, c := range conns {
			c.Close()
		}
	}()

	log.Printf("Total %d connections are established", len(conns))

	tts := time.Second
	if *connections > 100 {
		tts = time.Millisecond * 5
	}

	for {
		for i := 0; i < len(conns); i++ {
			time.Sleep(tts)
			conn := conns[i]

			conn.Write([]byte("hello world\r\n"))
		}
	}
}

func setLimit() {
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
}
