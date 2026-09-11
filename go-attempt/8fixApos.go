package main

import "strings"

// FixApostrophes ensures proper spacing for single quotes around words.
func FixApostrophes(text string) string {
	var result strings.Builder
	runes := []rune(text)
	i := 0
	for i < len(runes) {
		if runes[i] == '\'' {
			// Add opening quote
			result.WriteRune('\'')
			i++

			// Skip any spaces immediately after opening quote
			for i < len(runes) && runes[i] == ' ' {
				i++
			}

			// Collect the word(s) inside quotes
			for i < len(runes) && runes[i] != '\'' {
				result.WriteRune(runes[i])
				i++
			}

			// Skip any spaces before closing quote
			for i < len(runes) && runes[i] == ' ' {
				i++
			}

			// Add closing quote if it exists
			if i < len(runes) && runes[i] == '\'' {
				result.WriteRune('\'')
				i++
			}
			continue
		}

		result.WriteRune(runes[i])
		i++
	}

	return result.String()
}