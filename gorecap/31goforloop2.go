package main

import (
	"fmt"
)

func LoopOLoopa() {
	c := [2]string{"A", "The"}
	d := [3]string{"Man", "Husband", "Father"}
	e := [3]string{"Provider", "Protector", "Leader"}

	for i := 0; i < 2; i++ {
		for j := 0; j < len(d); j++ {
			if i == 1 {
				continue
			}
			fmt.Println(c[i], d[j])
		}
		fmt.Println()

		for k := 0; k < len(e); k++ {
			if i == 0 {
				continue
			}
			fmt.Println(c[i], e[k])
		}
	}
}
