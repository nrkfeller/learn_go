package main

import "fmt"

func main() {

	rw := BasketballPlayer{
		Athlete: Athlete{
			name:   "Russel Westbrook",
			age:    27,
			height: 200,
			weight: 190,
		},
		team:   "OKC",
		salary: 15,
	}

	fmt.Println(fullname(&rw))
	fmt.Println(rw.name)
	fmt.Println(rw.team)
}

// Athlete is the base struct
type Athlete struct {
	name   string
	age    int
	height int
	weight int
}

// BasketballPlayer inherits from athlete
type BasketballPlayer struct {
	Athlete
	team   string
	salary int
}

func fullname(bp *BasketballPlayer) string {
	return bp.name + bp.team
}
