package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Validate arguments
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		return
	}

	// Process input
	input := strings.ReplaceAll(os.Args[1], "\\n", "\n")

	// Here is our Default banner
	filename := "standard.txt"

	// Override the default banner if banner is provided
	if len(os.Args) == 3 {
		filename = os.Args[2] + ".txt"
	}

	// Read file
	lines, err := readBannerFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	//Convert our banner to a map
	asciiMap := buildAsciiMap(lines)

	// Render (must support multi-line)
	renderAscii(input, asciiMap)
}
