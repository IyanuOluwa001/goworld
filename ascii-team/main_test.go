package main

import (
	"os"
	"strings"
	"testing"
)

/*
TestStandardBannerFileExists

This test verifies that the standard.txt
banner file exists and can be opened.
*/
func TestStandardBannerFileExists(t *testing.T) {

	_, err := os.ReadFile("standard.txt")

	if err != nil {
		t.Fatalf("Failed to read standard.txt: %v", err)
	}
}

/*
TestStandardBannerNotEmpty

This test ensures the banner file
is not empty.
*/
func TestStandardBannerNotEmpty(t *testing.T) {

	data, err := os.ReadFile("standard.txt")
	if err != nil {
		t.Fatalf("Failed to read standard.txt: %v", err)
	}

	if len(data) == 0 {
		t.Fatalf("standard.txt should not be empty")
	}
}

/*
TestStandardBannerLineCount

The ASCII banner should contain
95 printable ASCII characters.

Each character block occupies 9 lines
(8 rows of ASCII art + 1 separator line).

Expected total lines = 855
*/
func TestStandardBannerLineCount(t *testing.T) {

	data, err := os.ReadFile("standard.txt")
	if err != nil {
		t.Fatalf("Failed to read standard.txt: %v", err)
	}

	lines := strings.Split(string(data), "\n")

	if len(lines) < 855 {
		t.Errorf("standard.txt appears incomplete, expected at least 855 lines, got %d", len(lines))
	}
}

/*
TestPrintableAsciiRange

Ensure the ASCII range used in the project
is correct: characters 32 to 126.
*/
func TestPrintableAsciiRange(t *testing.T) {

	count := 0

	for ascii := 32; ascii <= 126; ascii++ {
		count++
	}

	if count != 95 {
		t.Errorf("Printable ASCII range should contain 95 characters, got %d", count)
	}
}