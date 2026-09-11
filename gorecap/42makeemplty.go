package main

import (
	"fmt"
)

func mapIndo() {
	// make() actually creates and initializes the map
	// so memory for a is created already and its ready to use
	// It is not NIL
	a := make(map[string]string)

	// a map is declared but not initialized with make and no
	// memory allocated
	var b map[string]string

	fmt.Println(a == nil) // false
	fmt.Println(b == nil) // true

	// map key can be any of the following data types which == can be defined for.
	// Booleans, Numbers, Strings, Arrays, Pointers, Structs, Interfaces (as long as the dynamic type supports equality)

	// Invalid key type are: Slices, Maps, Functions
	// Map value can be any type
}
