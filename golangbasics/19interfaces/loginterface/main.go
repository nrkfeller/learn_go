package main

import "fmt"

func main() {
	cl := ConsoleLogger{}
	cl.Log("potato")

	al := AssLogger{}
	al.Log("assloging")
}

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (l ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

type AssLogger struct{}

func (al *AssLogger) Log(message string) {
	fmt.Printf("%v \n", message)
}
