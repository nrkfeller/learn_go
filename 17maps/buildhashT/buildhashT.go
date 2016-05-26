package main

import "fmt"

func main() {

	n := HashBucket("go", 12)
	fmt.Println(n)
}

// HashBucket takes in word and number of buckets and returns the bucket to place that word
func HashBucket(word string, buckets int) int {
	return int(word[0]) % buckets
}
