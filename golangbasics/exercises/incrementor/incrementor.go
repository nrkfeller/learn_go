package main

import "fmt"
import "strconv"

func main() {
	cout := inc(10)

	for n := range cout {
		fmt.Println(n)
	}
}

func inc(n int) chan string {
	out := make(chan string)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func(i int) {
			for j := 0; j < 20; j++ {
				out <- fmt.Sprint(strconv.Itoa(i) + " : " + strconv.Itoa(j))
			}
			done <- true
		}(i)
	}

	go func() {
		for i := 0; i < 2; i++ {
			<-done
		}
		close(out)
	}()
	return out
}
