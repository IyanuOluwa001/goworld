package main

import (
	"fmt"
	"strings"
)

func renderAscii(input string, asciiMap map[rune][]string) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			fmt.Println()
			continue
		}
		for row := 0; row < 8; row++ {
			for _, char := range line {
				block, ok := asciiMap[char]
				if !ok {
					continue
				}
				fmt.Print(block[row])
			}
			fmt.Println()
		}
	}
}
