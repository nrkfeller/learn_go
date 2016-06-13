package main

import "fmt"

func main() {
	prius := car{}
	x6 := car{}
	//cars := []car{prius, x6}

	b747 := plane{}
	cseries := plane{}
	//planes := []plane{b747, cseries}

	titanic := boat{}
	rowboat := boat{}
	//boats := []boat{titanic, rowboat}

	// this interfaces for all vehicles!
	rides := []vehicles{prius, x6, b747, cseries, titanic, rowboat}
	fmt.Println(rides)
	fmt.Println("12".(string))
}

type vehicles interface{}

type vehicle struct {
	Seats    int
	MaxSpeed int
	Color    string
}

type car struct {
	vehicle
	NumWheels int
}

type plane struct {
	vehicle
	MaxAlt int
}

type boat struct {
	vehicle
	Length int
}
