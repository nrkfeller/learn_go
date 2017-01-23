package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "nick feller"
	tpl :=
		`
  <!DOCTYPE html>
  <html lang = "en">
  <meta charset="UTF-8">
  <head>
  <title>Hello World</title>
  </head>
  <body>
  <h1>` + name + `</h1>
  </body>
  </html>
  `

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("bad news bears", err)
	}
	defer nf.Close()
	io.Copy(nf, strings.NewReader(tpl))
}
