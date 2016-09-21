package main

import "fmt"

func main() {
	startFactory(3, 8, 4, 4)
}

func startFactory(numcars, wdw, ew, pw int) {
	wheelWorker := make(chan int, wdw)
	engineWorker := make(chan int, ew)
	paintWorker := make(chan int, pw)

	for i := 0; i < wdw; i++ {
		wheelWorker <- i
	}
	for i := 0; i < ew; i++ {
		engineWorker <- i
	}
	for i := 0; i < pw; i++ {
		paintWorker <- i
	}

	cars := make(chan car, numcars)
	for i := 0; i < numcars; i++ {
		go makecar(cars, wheelWorker, engineWorker, paintWorker)
	}

	for n := range cars {
		fmt.Println(n)
	}
}

func makecar(cars chan car, wheelWorker chan int, engineWorker chan int, paintWorker chan int) {

	cars <- car{1, true}
}

type car struct {
	state    int
	complete bool
}
