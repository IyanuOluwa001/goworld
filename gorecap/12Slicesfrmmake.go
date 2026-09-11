// If we don't define capacity here, it will equal to length
// slice_name := make([]type, length, capacity)
package main

import (
	"fmt"
)

func SliceMaker() {
	array1 := []string{"Tola", "Timi", "Wole", "Bukky", "Iyanu"}
	slit := array1[2:4]
	slit1 := []int{0, 1, 2, 3}
	slitmake := make([]int, 5, 10)
	slitmake2 := make([]int, 5)

	fmt.Printf("array1 = %v\n", array1)
	fmt.Printf("slit = %v\n", slit)
	fmt.Printf("slit1 = %v\n", slit1)
	fmt.Printf("slitmake = %v\n", slitmake)
	fmt.Println()

	fmt.Printf("length = %d\n", len(slitmake))
	fmt.Printf("capacity = %d\n", cap(slitmake))
	fmt.Println()

	fmt.Printf("slit2 = %v\n", slitmake2)
	fmt.Printf("length =%d\n", len(slitmake2))
	fmt.Printf("capacity = %d\n", cap(slitmake2))
	fmt.Println()
}
