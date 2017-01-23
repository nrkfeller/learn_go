package main

import "fmt"

func main() {
	name := "Nick Feller"
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
	fmt.Println(tpl)
}
