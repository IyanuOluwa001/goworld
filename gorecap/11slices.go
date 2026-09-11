package main

import (
	"fmt"
)

func SlicePrac() {
	//[]datatype{values} e.g slice_name := []datatype{values}
	//create slice from an array
	//using the make() function
	//We use len() to return length of a slice
	//We use cap() to return capacity of a slice

	slice1 := []int{}
	slice2 := []int{1, 2, 3}
	slice3 := []string{"Go", "Slices", "Are", "Powerful"}
	arr1 := [6]int{1, 2, 3, 4, 5, 6}
	slice4 := arr1[2:4]
	slice5 := arr1[0:6]

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	fmt.Println(slice4)
	fmt.Printf("slice = %v\n", slice4)
	fmt.Printf("Length = %d\n", len(slice4))   // total length is 4-2=2
	fmt.Printf("capacity = %d\n", cap(slice4)) // total capacity is 6-2=4
	// i.e numbers of elements starting at index 2 is 12,13, 14, 15 which
	// is equal 4.
	fmt.Printf("Capacity =%d\n", cap(slice5))
}
