package main

import "fmt"

// Maps are references to hash tables
func mapRefere() {
	a := map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := a

	fmt.Println(a)
	fmt.Println(b)

	b["year"] = "1970"
	fmt.Println("After change to b:")

	fmt.Println(a)
	fmt.Println(b)
}
