package main

import (
	"strings"
	"unicode"
)

func FixPunctuation(tokens []string) []string {

	text := RebuildText(tokens)

	text = FixStandardPunctuation(text)

	text = FixApostrophes(text)

	return Tokenize(text)
}

func FixStandardPunctuation(text string) string {

	var result strings.Builder

	runes := []rune(text)

	i := 0

	for i < len(runes) {

		r := runes[i]

		if isPunctuation(r) {

			if result.Len() > 0 && result.String()[result.Len()-1] == ' ' {

				tmp := result.String()

				result.Reset()

				result.WriteString(tmp[:len(tmp)-1])
			}

			start := i

			for i < len(runes) && isPunctuation(runes[i]) {
				i++
			}

			group := string(runes[start:i])

			result.WriteString(group)

			if i < len(runes) && !unicode.IsSpace(runes[i]) {

				result.WriteRune(' ')
			}

			continue
		}

		result.WriteRune(r)

		i++
	}

	return strings.TrimSpace(result.String())
}

func FixApostrophes(text string) string {
	var result strings.Builder
	runes := []rune(text)
	i := 0
	for i < len(runes) {
		if runes[i] == '\'' {
			// Add opening quote
			result.WriteRune('\'')
			i++

			// Skip spaces immediately after opening quote
			for i < len(runes) && runes[i] == ' ' {
				i++
			}

			// Collect content inside quotes
			var inside strings.Builder
			for i < len(runes) && runes[i] != '\'' {
				inside.WriteRune(runes[i])
				i++
			}

			// Trim any trailing spaces before closing quote
			content := strings.TrimRight(inside.String(), " ")
			result.WriteString(content)

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

func isPunctuation(r rune) bool {

	return r == '.' ||
		r == ',' ||
		r == '!' ||
		r == '?' ||
		r == ':' ||
		r == ';'
}

func IsArticleA(word string) bool {

	return word == "a" || word == "A"
}

func StartsWithVowelOrH(word string) bool {

	if len(word) == 0 {
		return false
	}

	first := unicode.ToLower(rune(word[0]))

	return first == 'a' ||
		first == 'e' ||
		first == 'i' ||
		first == 'o' ||
		first == 'u' ||
		first == 'h'
}

func FixArticleCase(article string) string {

	if article == "A" {
		return "An"
	}

	return "an"
}
