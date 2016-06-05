package main

import (
	"bufio" // bucket for data
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	words := make(map[string]string)

	// create buffer and place the body of the results in there - returns scanner
	sc := bufio.NewScanner(res.Body)
	// split the scanner by words
	sc.Split(bufio.ScanWords)

	// goes through all the elements(in this case words) in the scanner
	for sc.Scan() {
		// place this is map
		words[sc.Text()] = ""
	}

	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	i := 0
	for k := range words {
		fmt.Println(k)
		if i < 300 {
			break
		}
		i++
	}

}
