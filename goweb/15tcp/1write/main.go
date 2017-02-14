package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "sup hello")
		fmt.Fprintln(conn, "sup to you too")
		fmt.Fprintf(conn, "doin good?")
		conn.Close()
	}

}
