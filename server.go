package main

import (
	"fmt"
	"io"
	"net"
)

func StartServer(c chan<- int) {
	listener, err := net.Listen("tcp", ":5555")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	// mark server as ready
	c <- 1
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept error: %+v\n", err.Error())
			continue
		}
		go handleConnection(conn)
	}
}

// handleConnection handles a single connection.
func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("read error: %+v\n", err.Error())
			break
		}
		recvContent := string(buf[:n])
		fmt.Printf("read %d bytes: %s\n", n, recvContent)
	}
}
