package main

// helps keep the scope tight!

// a is not accessible anywhere else

import "fmt"

func main() {

	mybool := true

	if a := "itsatruu"; mybool {
		fmt.Println(a)
	}
}

/*

if err := file.Chmod(0664); err != nil {
  fmt.Println("no errors")
}

*/
