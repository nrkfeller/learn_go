package main

import (
	"io"
	"net/http"
)

type DogHandler int
type CatHandler int

// handler can implement ServeHTTP and receives a ResponseWriter and *Request
func (h DogHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dowg")
}

func (h CatHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "catsss")
}

func main() {

	// created values of type Handler
	var dog DogHandler
	var cat CatHandler

	// we create a pointer to a serve mux
	mux := http.NewServeMux()

	// serve mux can call handler and pass in a string and a handler
	mux.Handle("/dog/", dog)
	mux.Handle("/cat/", cat)

	// serve up server and take in port and handler
	http.ListenAndServe(":9000", mux)
	// note that mux is a handler that has taken in handlers
}
