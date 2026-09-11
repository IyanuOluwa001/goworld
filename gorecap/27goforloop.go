package main

import (
	"fmt"
)

func GoFoLoopy() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println()

	for j := 0; j <= 100; j += 10 {
		fmt.Printf("This value = %v\n", j)
	}
}
