package main

import (
	"log"
	"net/http"
)

func main() {

	// log fatal is no necessary, but it logs the errors, could be useful!
	log.Fatal(http.ListenAndServe(":9000", http.FileServer(http.Dir("."))))
}
