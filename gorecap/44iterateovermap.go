package main

import (
	"fmt"
)

func iterateMap() {
	// Defining map1
	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	// Defining map2 with make - string/string
	f := make(map[string]string)
	f["brand"] = "Ford"
	f["model"] = "Mustang"
	f["year"] = "1964"
	f["day"] = ""

	// Defining map3 with make - stirng/int
	g := make(map[string]int)
	g["Mile 12"] = 1
	g["Ketu"] = 2
	g["Ojota"] = 3
	g["Fadeyi"] = 4

	// Defining slice
	e := []int{1, 3, 4, 2, 5}

	// Defining array
	b := [3]string{"My", "Business", "Space"}

	// Iterating over a map in specific order
	var c []string
	c = append(c, "one", "two", "three", "four")

	// Printing a map normally
	for k, v := range a {
		fmt.Printf("%v: %v,", k, v)
	}
	fmt.Println()
	fmt.Println(b)
	fmt.Println()
	fmt.Println(e)

	for _, element := range c {
		fmt.Printf("%v : %v, ", element, a[element])
	}
}
