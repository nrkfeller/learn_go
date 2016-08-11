package main

import (
	"flag"
	"fmt"

	"github.com/sunfmin/fanout"
)

var workers = flag.Int("workers", 60, "并发数")

func main() {

	words := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < 1000000; i++ {
		words = append(words, i)
	}

	inputs := []interface{}{}
	for _, word := range words {
		inputs = append(inputs, word)
	}

	results, err2 := fanout.ParallelRun(*workers, func(input interface{}) (interface{}, error) {
		word := input.(int)
		outr := domainAvailable(word)
		return outr, nil
	}, inputs)

	fmt.Println("Finished ", len(results), ", Error:", err2)

}
func domainAvailable(word int) (ch int) {

	ch = word * word
	return
}
