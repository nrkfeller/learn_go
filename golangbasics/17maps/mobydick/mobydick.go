package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	scanner.Split(bufio.ScanWords)
	buckets := make([]int, 12)

	for scanner.Scan() {
		n := HashBucket(scanner.Text())
		buckets[n]++
	}

	fmt.Println(buckets)

}

// HashBucket is the hash function
func HashBucket(word string) int {
	wordTot := 0
	for _, i := range word {
		wordTot += int(i)
	}
	return wordTot % 12
}
