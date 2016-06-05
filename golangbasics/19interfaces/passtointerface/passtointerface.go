package main

import "fmt"

func main() {

	breez := cat{"miaw"}
	noah := dog{"bark"}
	tweety := bird{"churp"}

	// watch out! be careful with empty interfaces
	allanims := []interface{}{breez, noah, tweety}

	fmt.Println(allanims)

}

type cat struct {
	noise string
}

type dog struct {
	noise string
}

type bird struct {
	noise string
}
