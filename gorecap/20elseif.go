package main

import (
	"fmt"
)

func Elifa() {
	time := 20
	if time < 10 {
		fmt.Println("Good morning")
	} else if time < 20 {
		fmt.Println("Good Afternoon")
	} else {
		fmt.Println("Good evening")
	}
}
