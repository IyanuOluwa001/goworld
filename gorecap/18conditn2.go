package main

import (
	"fmt"
)

func ConDit() {
	time := 20
	if time < 18 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}
}

func conDiTions() {
	age := 20
	expertise := 2

	if age == 20 {
		fmt.Println("Not eligible")
	} else if age > 20 && expertise >= 2 {
		fmt.Println("Qualified")
	}
}
