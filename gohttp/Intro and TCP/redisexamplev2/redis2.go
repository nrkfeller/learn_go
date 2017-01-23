package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func redisServer(commands chan Command) {
	var data = make(map[string]string)
	for cmd := range commands {
		if len(cmd.Fields) < 2 {
			cmd.Result <- "Expects at least 2 args"
			continue
		}
		fmt.Println("Got Command", cmd)

		switch cmd.Fields[0] {
		case "GET":
			if len(cmd.Fields) != 2 {
				cmd.Result <- "Invalid number of arguments"
				continue
			}
			_, prs := data[cmd.Fields[1]]
			if prs {
				//fmt.Println(cmd.Fields[0], "has value", data[cmd.Fields[0]])
				cmd.Result <- data[cmd.Fields[1]]
			} else {
				cmd.Result <- cmd.Fields[1] + " is not present!"
			}
		case "SET":
			if len(cmd.Fields) != 3 {
				cmd.Result <- "Invalid number of arguments"
				continue
			}
			key := cmd.Fields[1]
			data[key] = cmd.Fields[2]
			cmd.Result <- "Setting " + key + " to " + data[key]
		case "DEL":
			if len(cmd.Fields) != 2 {
				cmd.Result <- "Invalid number of arguments"
				continue
			}
			_, prs := data[cmd.Fields[1]]
			if prs {
				cmd.Result <- "deleting : " + cmd.Fields[1]
				delete(data, cmd.Fields[1])
				cmd.Result <- ""
			} else {
				cmd.Result <- cmd.Fields[1] + " is not present!"
			}
		default:
			//io.WriteString(conn, "INVALID COMMAND : "+fs[0]+"\n")
			cmd.Result <- "INVALID COMMAND : " + cmd.Fields[0] + "\n"
		}
	}
}

func handle(commands chan Command, conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)
		result := make(chan string)
		commands <- Command{
			Fields: fs,
			Result: result,
		}
		io.WriteString(conn, <-result+"\n")
	}
}

// Command is typed and passed into the channel
type Command struct {
	Fields []string
	Result chan string
}

func main() {
	li, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	commands := make(chan Command)
	go redisServer(commands)

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		// our big problem, only one connection at a time!
		go handle(commands, conn)
	}
}
