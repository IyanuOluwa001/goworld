package main

import (
	"os"
	"testing"
)

func TestAsciiArt(t *testing.T) {

	tests := []struct {
		name         string
		input        string
		banner       string
		expectedFile string
	}{
		{
			name:         "single valid input",
			input:        "A",
			banner:       "standard",
			expectedFile: "testdata/single_A.txt",
		},
		{
			name:         "empty input",
			input:        "",
			banner:       "standard",
			expectedFile: "testdata/empty.txt",
		},
		{
			name:         "new line",
			input:        "\\n",
			banner:       "standard",
			expectedFile: "testdata/newline.txt",
		},
	}

	for _, tt := range tests {

		results, err := asciiArt(tt.input, tt.banner)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		data, err := os.ReadFile(tt.expectedFile)

		if err != nil {
			t.Fatalf("could not read expected file: %v", err)
		}

		expected := string(data)

		if results != expected {
			t.Errorf(
				"\nTEST: %s\nGOT:\n%q\n\nWANT:\n%q\n",
				tt.name,
				results,
				expected,
			)
		}
	}
}
