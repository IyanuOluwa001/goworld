package main

import (
	"strings"
	"unicode"
)

func FixStandardPunctuation(text string) string {

	var result strings.Builder
	runes := []rune(text)
	i := 0

	for i < len(runes) {
		r := runes[i]

		if isPunctuation(r) {
			// Remove preceding space
			if result.Len() > 0 && result.String()[result.Len()-1] == ' ' {
				tmp := result.String()
				result.Reset()
				result.WriteString(tmp[:len(tmp)-1])
			}

			// Collect group of punctuations
			start := i
			for i < len(runes) && isPunctuation(runes[i]) {
				i++
			}
			group := string(runes[start:i])
			result.WriteString(group)

			// Decide whether to add space after
			if i < len(runes) {
				// Don't add space if next rune is punctuation
				if !isPunctuation(runes[i]) && !unicode.IsSpace(runes[i]) {
					result.WriteRune(' ')
				} else if unicode.IsSpace(runes[i]) {
					result.WriteRune(' ')
				}
			}

			continue
		}

		result.WriteRune(r)
		i++
	}

	return strings.TrimSpace(result.String())
}

func isPunctuation(r rune) bool {
	return r == '.' || r == ',' || r == '!' || r == '?' || r == ':' || r == ';'
}

/*
package main

import (func IsModifier(word string) bool {
	return len(word) > 2 && word[0] == '(' && word[len(word)-1] == ')'
}
	"strings"
)

func FixStandardPunctuation(text string) string {

	var result strings.Builder
	runes := []rune(text)

	i := 0
	for i < len(runes) {

		if isPunctuation(runes[i]) {

			// Remove space before punctuation/group
			if result.Len() > 0 {
				str := result.String()
				if str[len(str)-1] == ' ' {
					result.Reset()
					result.WriteString(str[:len(str)-1])
				}
			}

			// Collect full punctuation group
			start := i
			for i < len(runes) && isPunctuation(runes[i]) {
				i++
			}

			group := string(runes[start:i])
			result.WriteString(group)

			// Add space after group if needed
			if i < len(runes) && runes[i] != ' ' {
				result.WriteRune(' ')
			}

			continue
		}

		result.WriteRune(runes[i])
		i++
	}

	return strings.TrimSpace(result.String())
}

func isPunctuation(r rune) bool {
	return r == '.' ||
		r == ',' ||
		r == '!' ||
		r == '?' ||
		r == ':' ||
		r == ';'
}
*/

/*
func FixStandardPunctuation(text string) string {

	var result strings.Builder
	runes := []rune(text)

	for i := 0; i < len(runes); i++ {

		current := runes[i]

		if isPunctuation(current) {

			// If next character is also punctuation → grouped punctuation
			if i+1 < len(runes) && isPunctuation(runes[i+1]) {
				result.WriteRune(current)
				continue
			}

			// Remove space before punctuation
			if result.Len() > 0 {
				last := result.String()[result.Len()-1]
				if last == ' ' {
					str := result.String()
					result.Reset()
					result.WriteString(str[:len(str)-1])
				}
			}

			result.WriteRune(current)

			// Add space after if next character is not space
			if i+1 < len(runes) && runes[i+1] != ' ' {
				result.WriteRune(' ')
			}

		} else {
			result.WriteRune(current)
		}
	}

	return strings.TrimSpace(result.String())
}

func isPunctuation(r rune) bool {
	return r == '.' ||
		r == ',' ||
		r == '!' ||
		r == '?' ||
		r == ':' ||
		r == ';'
}
*/

/*
func FixStandardPunctuation(text string) string {

	punctuations := []string{".", ",", "!", "?", ":", ";"}

	for _, p := range punctuations {

		// remove space before punctuation
		text = strings.ReplaceAll(text, " "+p, p)

		// ensure one space after if needed
		text = strings.ReplaceAll(text, p, p+" ")
	}

	// remove double spaces created
	text = strings.ReplaceAll(text, "  ", " ")

	// trim trailing space
	text = strings.TrimSpace(text)

	return text
}
	*/