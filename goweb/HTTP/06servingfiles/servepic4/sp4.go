package main

import (
	"io"
	"net/http"
	"strings"
)

func main() {

	// this servers all the files in a directory!
	// you can even look at sp4.go
	// http.Handle("/", http.FileServer(http.Dir(".")))

	// Serves files from a specific directory and places it in /pictures
	http.Handle("/pictures/", http.StripPrefix("/pictures", http.FileServer(http.Dir("./assets"))))

	http.HandleFunc("/food/", food)

	http.ListenAndServe(":9000", nil)
}

func food(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html ; charset=utf-8")
	var name string

	fs := strings.Split(req.URL.Path, "/")

	if len(fs) >= 3 {
		name = fs[2]
	}

	io.WriteString(res, `
    Apple Type: `+name+` </br>
    <img src="/pictures/pic.jpg ">
    `)
}
