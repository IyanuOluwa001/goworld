package main

import (
	"fmt"
)

func LoopOLoop() {
	a := "dodo"
	b := "mayana"
	c := [2]string{"A", "The"}
	d := [3]string{"Man", "Husband", "Father"}
	e := [3]string{"Provider", "Protector", "Leader"}

	fmt.Println(a + " " + b)
	fmt.Println(a, b)
	fmt.Println()

	for i := 0; i < len(c); i++ {
		if i == 1 {
			continue
		}
		for j := 0; j < len(d); j++ {
			fmt.Println(c[i], d[j])
		}
		fmt.Println()

		for i := 0; i < len(c); i++ {
			for k := 0; k < len(e); k++ {
				fmt.Println(c[i], e[k])
			}
			fmt.Println()
		}
	}
}
