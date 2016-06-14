package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// resp here is a response, and resp.Body is an io.Reader
	resp, _ := http.Get(os.Args[1])
	fmt.Fprintln(os.Stderr, resp)
}
