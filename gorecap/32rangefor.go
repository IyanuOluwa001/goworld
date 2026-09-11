package main

import (
	"fmt"
)

func Ranger() {
	fruits := [3]string{"apple", "orange", "banana"}
	//Print Index and value
	/*
		// idx stores the index, val stores the value
		for idx, val := range fruits{
			fmt.Printf("%v\t%v\n", idx, val)
		}
	*/

	// Print the value of the array excluding the index
	// for _, val := range fruits{
	//	fmt.Printf("%v\n",val)
	for idx := range fruits {
		fmt.Printf("%v\n", idx)
	}
}
