package main

import (
	"fmt"
)

func AssignIt() {
	x := 5
	fmt.Printf("x is %b\n", x)
	x <<= 3 // x=x<<3
	fmt.Printf("x is now %03b\n", x)
	x >>= 4 // x=x<<4
	fmt.Printf("x now becomes %03b\n", x)
}
