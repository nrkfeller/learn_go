package main

import (
	"fmt"
	"strconv"
)

func makeCakeAndSend(cs chan string, cakes []string, done chan bool) {

	for i := 0; i < len(cakes); i++ {
		cakeName := cakes[i] + " Cake " + strconv.Itoa(i)
		cs <- cakeName //send a strawberry cake
	}
	done <- true
}

func receiveCakeAndPack(cs chan string, done chan bool) {
	for s := range cs {
		fmt.Println("Packing received cake:", s)
	}
	done <- true
}

func main() {

	done := make(chan bool)
	cs := make(chan string)
	types := []string{"rasb", "strawb", "apple", "tofu", "peach"}
	go makeCakeAndSend(cs, types, done)
	go receiveCakeAndPack(cs, done)
	<-done
	<-done

}
