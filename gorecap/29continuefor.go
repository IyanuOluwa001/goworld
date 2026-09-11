package main

import (
	"fmt"
)

func Contina() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			// continue skips the round,
			// so 3 would be skipped here
			continue
		}
		fmt.Println(i)
	}
	fmt.Println()

	for i := 4; i < 20; i += 2 {
		if i == 10 {
			break
		}
	}
}
