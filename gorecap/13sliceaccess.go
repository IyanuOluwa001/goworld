package main

import (
	"fmt"
)

func Slico() {
	prices := []int{10, 20, 30}

	// To Append elements to a slice
	// slice_name = append(slice_name, element1, element2, ...)
	myslice1 := []int{1, 2, 3, 4, 5, 6}

	// Change elements of an array for strings
	// cars[0] = "Toyota"

	// Access first element by printing
	fmt.Println(prices[0])
	fmt.Println()

	fmt.Println(prices)
	prices[1] = 40 // Change element of a slice for Int
	fmt.Println(prices)
	fmt.Println()

	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))
	fmt.Println()

	// Appending elements to a slice
	myslice2 := []int{7, 8, 9, 10, 11, 12}
	myslice2 = append(myslice2, 20, 21)
	fmt.Printf("myslice2 = %v\n", myslice2)
	fmt.Printf("length = %d\n", len(myslice2))
	fmt.Printf("capacity = %d\n", cap(myslice2))
	fmt.Println()

	// Append one slice to another
	myslice3 := append(myslice1, myslice2...)
	fmt.Printf("myslice2 = %v\n", myslice3)
	fmt.Printf("length = %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3))
	fmt.Println()

	// Change the length of a Slice
	arr1 := [6]int{13, 14, 15, 16, 17, 18}
	myslice4 := arr1[1:5]
	fmt.Printf("Slice4 = %v\n", myslice4)
	fmt.Printf("length = %d\n", len(myslice4))
	fmt.Printf("capacity = %d\n", cap(myslice4))
	fmt.Println()

	// Change length by reslicing the array
	myslice4 = arr1[1:3]
	fmt.Printf("myslice4 = %v\n", myslice4)
	fmt.Printf("length = %d\n", len(myslice4))
	fmt.Printf("capacity = %d\n", cap(myslice4))
	fmt.Println()

	// Change length by appending items
	myslice4 = append(myslice4, 22, 23)
	fmt.Printf("myslice4 = %v\n", myslice4)
	fmt.Printf("length = %d\n", len(myslice4))
	fmt.Printf("capacity = %d\n", cap(myslice4))

	//Copy() function
	/*
		The function creates a new underlying array with only the
		required elements for the slice. This will reduce the memory
		for the program (this is done becuase go loads the underlying
		elements into memory when using slices. so if the array is large
		and you need only a few elements, it is better to copy the elements
		using the copy() function)
	*/

	// Original slice
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity =%d\n", cap(numbers))
	fmt.Println()

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))
}
