package main

import (
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatalln(err)
	}
	li.Close()
}
