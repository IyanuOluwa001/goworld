package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Ensure exactly one argument is provided
	if len(os.Args) != 2 {
		fmt.Println("Error: please provide a single text argument")
		return
	}

	input := os.Args[1]
	// handle a next line string
	if input == "\n" || input == "\\n" {
		fmt.Println()
		return
	}
	AsciiArt(input)
}

// AsciiArt prints the input string as ASCII art using standard.txt
func AsciiArt(input string) {
	// Read the ASCII font file
	content, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading standard.txt:", err)
		return
	}

	lines := strings.Split(string(content), "\n")
	userLines := strings.Split(input, "\\n") // Handle multi-line input

	for _, line := range userLines {
		if line == "" { // Empty line
			fmt.Println()
			continue
		}

		// Each ASCII character is 8 rows tall
		for row := 0; row < 8; row++ {
			for _, char := range line {
				// Only process printable ASCII characters
				if char < 32 || char > 126 {
					continue
				}
				// Map character to its position in the font file
				idx := (int(char)-32)*9 + 1
				fmt.Print(lines[idx+row])
			}
			fmt.Println()
		}
	}
}
