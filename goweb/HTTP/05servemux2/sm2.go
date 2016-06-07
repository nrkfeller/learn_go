package main

import (
	"io"
	"net/http"
	"strings"
)

func main() {
	var captain CapHandler
	var firstmate FMHandler

	mux := http.NewServeMux()
	mux.Handle("/cap/", captain)
	mux.Handle("/fm/", firstmate)

	http.ListenAndServe(":9000", mux)
}

type CapHandler int
type FMHandler int

func (h FMHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html ; charset=utf-8")
	var fmname string
	fs := strings.Split(req.URL.Path, "/")
	if len(fs) >= 3 {
		fmname = fs[2]
	}
	io.WriteString(res, "Hello First Mate : "+fmname)
}

func (h CapHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html ; charset=utf-8")
	var capname string
	fs := strings.Split(req.URL.Path, "/")
	if len(fs) >= 3 {
		capname = fs[2]
	}
	io.WriteString(res, "Hello Captain : "+capname)
}
