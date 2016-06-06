/*
Just a simple KV store with
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

var data = make(map[string]string)

func handle(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)
		if len(fs) < 1 {
			continue
		}
		switch fs[0] {
		case "GET":
			get(fs[1:])
		case "SET":
			set(fs[1:])
		case "DEL":
			del(fs[1:])
		default:
			io.WriteString(conn, "INVALID COMMAND : "+fs[0]+"\n")
		}
	}
}

func set(items []string) {
	if len(items) != 2 {
		fmt.Println("Invalid number of arguments")
		return
	}
	key := items[0]
	data[key] = items[1]
	fmt.Println("Setting", key, "to", data[key])
}

func get(items []string) {
	if len(items) != 1 {
		fmt.Println("Invalid number of arguments")
		return
	}
	_, prs := data[items[0]]
	if prs {
		fmt.Println(items[0], "has value", data[items[0]])
	} else {
		fmt.Println(items[0], "is not present!")
	}
}

func del(items []string) {
	if len(items) != 1 {
		fmt.Println("Invalid number of arguments")
		return
	}
	_, prs := data[items[0]]
	if prs {
		fmt.Println("deleting : ", items[0])
		delete(data, items[0])
	} else {
		fmt.Println(items[0], "is not present!")
	}
}

func main() {
	li, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		handle(conn)
	}
}
