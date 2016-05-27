package main

import (
	"fmt"
	"sort"
)

func main() {
	names := []string{"nick", "tu", "abe", "jin"}

	sort.Strings(names)

	fmt.Println(names)
}
