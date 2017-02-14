package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", hi)
	http.ListenAndServe(":8080", nil)
}

func hi(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hihihii")
}
