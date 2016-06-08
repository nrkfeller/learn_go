package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/asian/", asian)
	http.HandleFunc("/swiss/", swiss)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func asian(res http.ResponseWriter, req *http.Request) {
	fs := strings.Split(req.URL.Path, "/")
	io.WriteString(res, fs[len(fs)-1]+" is asian")
}

func swiss(res http.ResponseWriter, req *http.Request) {
	fs := strings.Split(req.URL.Path, "/")
	io.WriteString(res, fs[len(fs)-1]+" is swiss")
}
