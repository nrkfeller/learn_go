package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi papa Ilter m+ deployed")
}

// docker build -t hellotest .
// docker run -it --rm --name ma-instance -p 8080:8080 hellotest
