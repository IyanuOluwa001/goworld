package main

import (
	"strings"
	"testing"
)

func captureRenderAscii(input string, asciiMap map[rune][]string) string {
	var output strings.Builder
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		for row := 0; row < 8; row++ {
			for _, char := range line {
				block, ok := asciiMap[char]
				if !ok {
					continue
				}
				output.WriteString(block[row])
			}
			output.WriteString("\n")
		}
	}
	return output.String()
}
func TestBuildiAsciiMap(t *testing.T) {

	lines := []string{
		"  A  ",
		" A A ",
		"AAAAA",
		"A   A",
		"A   A",
		"A   A",
		"A   A",
		"     ",
	}
	asciiMap := buildAsciiMap(lines)
	if len(asciiMap) == 0 {
		t.Error("Expected error due to incomplete file, got nil")
	}
}

func TestRenderAscii_SingleChar(t *testing.T) {
	lines := []string{
		"  A  ",
		" A A ",
		"AAAAA",
		"A   A",
		"A   A",
		"A   A",
		"A   A",
		"     ",
	}

	asciiMap := map[rune][]string{
		'A': lines,
	}

	expected := "" +
		"  A  \n" +
		" A A \n" +
		"AAAAA\n" +
		"A   A\n" +
		"A   A\n" +
		"A   A\n" +
		"A   A\n" +
		"     \n"

	output := captureRenderAscii("A", asciiMap)
	if output != expected {
		t.Errorf("Got:\n%s\nExpected:\n%s", output, expected)
	}
}

func TestRenderAscii_MultiLine(t *testing.T) {
	linesA := []string{
		"  A  ", " A A ", "AAAAA", "A   A", "A   A", "A   A", "A   A", "     ",
	}

	asciiMap := map[rune][]string{
		'A': linesA,
	}

	input := "A\nA"
	output := captureRenderAscii(input, asciiMap)

	expected := "" +
		"  A  \n" +
		" A A \n" +
		"AAAAA\n" +
		"A   A\n" +
		"A   A\n" +
		"A   A\n" +
		"A   A\n" +
		"     \n" +
		"  A  \n" +
		" A A \n" +
		"AAAAA\n" +
		"A   A\n" +
		"A   A\n" +
		"A   A\n" +
		"A   A\n" +
		"     \n"

	if output != expected {
		t.Errorf("Got:\n%s\nExpected:\n%s", output, expected)
	}
}
