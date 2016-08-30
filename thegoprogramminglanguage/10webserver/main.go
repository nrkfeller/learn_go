package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	url := strings.Split(r.URL.Path, "/")
	fmt.Fprintf(w, "URL.Path = %q\n", url[len(url)-1])
}
