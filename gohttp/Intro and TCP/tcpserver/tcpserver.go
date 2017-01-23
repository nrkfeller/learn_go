/*
go run tcpserver.go -- navigate to localhost:9000

telnet is a TCP server that listens
of type : telnet localhost 9000 in terminal
*/

package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {

	// Listen returns "Type Listener"
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	// ln.Accept returns "a connection and an error"
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		// io.WriteString takes (a writer and a string) returns (and int and an error) -- but in this case we throw away both
		io.WriteString(conn, fmt.Sprint("Hello world \n", time.Now(), "\n"))
		conn.Close()
	}
}
