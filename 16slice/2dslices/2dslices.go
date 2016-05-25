package main

import "fmt"

func main() {

	var records [][]string

	// student 1
	student1 := make([]string, 4)
	student1[0] = "Feller"
	student1[1] = "Nicolas"
	student1[2] = "100.00"
	student1[3] = "99.99"
	records = append(records, student1)
	// student 1
	student2 := make([]string, 4)
	student2[0] = "Solomon"
	student2[1] = "Abel"
	student2[2] = "90.00"
	student2[3] = "79.99"
	records = append(records, student2)

	fmt.Println(records)

}
