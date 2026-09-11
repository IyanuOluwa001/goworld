package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// Program expects only one argument: the text to render
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . \"text\"")
		return
	}

	// Read banner file
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Split banner file into lines
	lines := strings.Split(string(data), "\n")
	// Replace "\n" with actual newline characters
	input := strings.ReplaceAll(os.Args[1], "\\n", "\n")

	// Build ASCII character map
	asciiMap := make(map[rune][]string)

	for ascii := 32; ascii <= 126; ascii++ {

		index := ascii - 32
		start := index * 9

		var block []string

		for i := 0; i < 8; i++ {
			block = append(block, lines[start+i])
		}

		asciiMap[rune(ascii)] = block
	}

	// Split input to support multiple lines
	inputLines := strings.Split(input, "\n")

	for _, line := range inputLines {

		// Preserve empty lines
		if line == "" {
			fmt.Println()
			continue
		}

		// Each ASCII character prints over 8 rows
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