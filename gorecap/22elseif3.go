package main

import (
	"fmt"
)

func FirstOff() {
	x := 30
	if x >= 10 {
		fmt.Println("x is greater or equal to 10")
	} else if x > 20 {
		fmt.Println("x is greater than 20")
	} else {
		fmt.Println("x is less than 10")
	}
}
