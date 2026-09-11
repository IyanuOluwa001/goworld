package main

import (
	"fmt"
)

func ComArat() {
	// The return value of a comparison is either true or false
	x := 5
	y := 3
	fmt.Println(x > y)  // returns true
	fmt.Println(x == y) // is x equal to y
	fmt.Println(x != y) // is x not equal y
	fmt.Println(x < y)  // is x less than y
	fmt.Println(x >= y) // is x greater than or equal to y
	fmt.Println(x <= y) // is x less than or equal to y
	fmt.Println()

	// Logical AND(&&): returns true if both statements are true
	fmt.Println(x < 5 && x < 10)

	// Logical OR(||): returns true if one of the statements is true
	fmt.Println(x < 5 || x < 4)

	// Logical NOT(!(s)): reverses the result, returns false if true, and true if false
	fmt.Println(!(x < 5 && x < 10))
	fmt.Println()

	// AND: Sets each bit to 1 if both bits are 1
	fmt.Printf("x = %b\n", x)
	fmt.Printf("y = %b\n", y)
	fmt.Printf("x & y is %b\n", x&y)

	// OR: Sets each bit to 1 if one of the two bits is 1
	fmt.Printf("x | y is %b\n", x|y)

	// XOR: Sets each bit to 1 if only one of two bits is 1
	fmt.Printf("x ^ y is %b\n", x^y)

	//<<: Zero fill left shift(shift left by pushing zeros in from the right)
	fmt.Printf("x << 2 is %b\n", x<<2)
	//>>: Signed right shift (shift right by pushing copies of the leftmost bit in
	//from the left, and let the rightmost bits fall off)
	fmt.Printf("x <<2 is %b\n", x<<2)
}
