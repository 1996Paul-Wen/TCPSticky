package main

import (
	"fmt"
	"net"
)

func StartClient(c <-chan int) {
	// wait for server to start
	<-c

	conn, err := net.Dial("tcp", ":5555")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := "Hello! How are you?"
		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Printf("write error: %+v\n", err.Error())
			break
		}
	}
}
