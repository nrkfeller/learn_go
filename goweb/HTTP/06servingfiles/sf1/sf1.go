package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(req.URL)
	})

	http.HandleFunc("/apple/", apple)
	http.HandleFunc("/pic.jpg", applepic)
	http.ListenAndServe(":9000", nil)
}

// file needs to be opened and copied
func applepic(res http.ResponseWriter, req *http.Request) {
	f, err := os.Open("pic.jpg")
	if err != nil {
		http.Error(res, "file not found", 404)
		return
	}
	defer f.Close()

	// Copy takes in destination and source!
	//                writer          reader
	io.Copy(res, f)
}

func apple(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html ; charset=utf-8")
	var appletype string
	fs := strings.Split(req.URL.Path, "/")

	if len(fs) >= 3 {
		appletype = fs[2]
	}
	io.WriteString(res, `
    Apple Type: `+appletype+` </br>
    <img src="/pic.jpg ">
    `)
}
