package main

import "fmt"

// func pinger(c chan string) {
// 	for i := 0; ; i++ {
// 		c <- "ping" + string(i)
// 	}
// }
//
// func printer(c chan string) {
// 	for {
// 		msg := <-c
// 		fmt.Println(msg)
// 		time.Sleep(time.Second * 1)
// 	}
// }

func main() {
	c := make(chan string)
	go pinger(c)
	go printer(c)

	var input string
	fmt.Scan(&input)
}
