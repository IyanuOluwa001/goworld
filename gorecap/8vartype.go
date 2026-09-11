package main

import (
	"fmt"
)

const Pi = 3.14

func VarVary() {
	var IsCat string = "Hello"
	var j int = 15
	// Constant cannot be declared with :=
	const World = "Nigeria"

	fmt.Printf("Iscat has value: %v and type: %T\n", IsCat, IsCat)
	fmt.Printf("j has value: %v and type: %T\n", j, j)
	fmt.Printf("j has value: %#v and type: %%\n", j)
	fmt.Printf("Iscat has value: %#v and type: %%\n", IsCat)
	fmt.Println(IsCat, World)
}
