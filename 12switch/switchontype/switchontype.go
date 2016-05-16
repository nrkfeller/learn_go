package main

import "fmt"

func main() {

	x := 12.234
	whattype(x)

}

func whattype(x interface{}) {
	var ret string
	switch x.(type) {
	case int, float64:
		ret = "number"
	case string:
		ret = "string"
	case bool:
		ret = "bool"
	}

	fmt.Println(ret)
}
