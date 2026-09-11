package main

import (
	"fmt"
)

func Ticker() {
	num := 20
	if num >= 10 {
		fmt.Println("This number pass 10 o")
		if num > 15 {
			fmt.Println("This number pass 15 seh!")
		}
	} else {
		fmt.Println("Num is less than 10")
	}
}
