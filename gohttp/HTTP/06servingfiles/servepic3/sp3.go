package main

import (
	"io"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/fruit/", apple)
	http.HandleFunc("/pic.jpg", applepic)

	http.ListenAndServe(":9000", nil)
}

func apple(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "text/html ; charset=utf-8")
	var fruitname string
	fs := strings.Split(req.URL.Path, "/")

	if len(fs) >= 3 {
		fruitname = fs[2]
	}

	io.WriteString(res, `
    Apple Type: `+fruitname+` </br>
    <img src="/pic.jpg ">
    `)
}

// best way for pictures!
func applepic(res http.ResponseWriter, req *http.Request) {
	// serves from current directory http.ServeFile(res, req, "img/pic.jpg") for another directory
	http.ServeFile(res, req, "pic.jpg")
	// this wants
}
