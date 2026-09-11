package main

import (
	"os"
	"strings"
	"testing"
)

func TestAsciiArt(t *testing.T) {

	// Create a builder to efficiently build the string
	var builder strings.Builder

	// Create fake ASCII banner content
	for i := 0; i < 95; i++ {
		for j := 0; j < 9; j++ {
			builder.WriteString("AAAA\n")
		}
	}

	// Convert builder to string
	content := builder.String()

	// Write fake banner file
	os.WriteFile("standard.txt", []byte(content), 0644)

	// Remove the file after test finishes
	defer os.Remove("standard.txt")

	// Run tests
	AsciiArt("A")
	AsciiArt("AB")
	AsciiArt("A\\nB")

	AsciiArt("hello")
	AsciiArt("HELLO")
	AsciiArt("HeLlo HuMaN")
	AsciiArt("1Hello 2There")
	AsciiArt("Hello\nThere")
	AsciiArt("Hello\n\nThere")
	AsciiArt("{Hello & There #}")
	AsciiArt("hello There 1 to 2!")
	AsciiArt("MaD3IrA&LiSboN")
	
}