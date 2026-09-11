package main

import (
	"fmt"
)

func makeMap() {
	// make() function is the right way to create an empty map.
	// creating an empty map in a different way and writing to it will cause a runtime panic.
	a := make(map[string]string) // The map is empty now
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"
	a["day"] = ""

	b := make(map[string]int)
	b["Mile 12"] = 1
	b["Ketu"] = 2
	b["Ojota"] = 3
	b["Fadeyi"] = 4

	c := map[string]string{"username": "tickerman", "category": "Investor", "level": "5", "year": "2022"}

	// Printmap
	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
	fmt.Println()

	// Access Map Elements
	fmt.Printf(a["brand"])
	fmt.Println()

	// Print plain map
	fmt.Println(a)
	fmt.Println()

	// Update map
	a["year"] = "1970"   // Updating an element
	a["color"] = "red"   // Adding an element
	a["gender"] = "male" // Adding an element
	fmt.Println(a)
	fmt.Println()

	delete(a, "gender")
	fmt.Println(a)
	fmt.Println()

	// Checking for the existing key and its value
	val1, ok1 := a["brand"]
	// Checking for the non-existing key and its value
	val2, ok2 := a["gender"]
	// Checking for existing key and its value
	val3, ok3 := a["year"]
	// Only checking for existing key and not its value
	_, ok4 := a["model"]
	val5, ok5 := a["day"]

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)
	fmt.Println(val5, ok5)
	fmt.Println()

	fmt.Println(c)
}
