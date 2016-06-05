package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.1234nicolasfeller.com/")
	if err != nil {
		log.Fatal(err)
	}

	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", page)
	fmt.Printf("%T", page)
}
