/*
different style. seems kind of simpler!
*/

package main

import (
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// HandleFunc takes a HandlerFunc!
	mux.HandleFunc("/OKC/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "GO OKCOKCOKC")
	})

	mux.HandleFunc("/GSW/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "BOOOO GSW")
	})

	http.ListenAndServe(":9000", mux)
}
