package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		handle(err)
		_, err = io.Copy(os.Stdout, resp.Body)
		// resp.Body.Close()
		handle(err)
		// fmt.Printf("%s", b)
		fmt.Println(resp.Status)
	}
}

func handle(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v", err)
		os.Exit(1)
	}
}
