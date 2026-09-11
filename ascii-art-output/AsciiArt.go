package main

import (
	"os"
	"strings"
)

func asciiArt(input, banner string) (string, error) {

	if input == "" {
		return "", nil
	}
	if input == "\\n" {
		return "\n", nil
	}

	data, err := os.ReadFile(banner + ".txt")
	if err != nil {
		return "", err
	}

	text := strings.Split(string(data), "\n")
	word := strings.Split(input, "\\n")

	var result strings.Builder

	for _, words := range word {

		if words == "" {
			result.WriteString("\n")
			continue
		}

		for i := 0; i < 8; i++ {
			for _, character := range words {
				index := (int(character)-32)*9 + 1 + i
				result.WriteString(text[index])
			}
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}
