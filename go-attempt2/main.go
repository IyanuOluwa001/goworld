package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// Ensure a filename argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <inputfile>")
		return
	}

	inputFile := os.Args[1]

	// Read file content
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	text := string(content)

	// Apply transformations
	text = applyModifiers(text)
	text = fixPunctuation(text)
	text = fixQuotes(text)
	text = fixAAn(text)

	// Write modified content to output.txt
	err = os.WriteFile("output.txt", []byte(text), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Processing complete. Output saved to output.txt")
}

//  Apply All Text Modifiers
func applyModifiers(text string) string {

	// Regex to match modifiers like:
	// (hex), (bin), (up), (low), (cap)
	// and also with numbers like (hex, 3)
	re := regexp.MustCompile(`\((hex|bin|up|low|cap)(,\s*\d+)?\)`)

	for {
		loc := re.FindStringIndex(text)
		if loc == nil {
			break
		}

		match := text[loc[0]:loc[1]]

		// Extract modifier and optional number
		parts := strings.Trim(match, "()")
		split := strings.Split(parts, ",")

		modifier := strings.TrimSpace(split[0])
		count := 1 // default is 1 word

		if len(split) > 1 {
			n, _ := strconv.Atoi(strings.TrimSpace(split[1]))
			count = n
		}

		// Split text before modifier into words
		before := strings.TrimSpace(text[:loc[0]])
		after := text[loc[1]:]

		words := strings.Fields(before)

		// Ensure enough words exist
		if len(words) < count {
			text = text[:loc[0]] + after
			continue
		}

		// Apply modification to last 'count' words
		for i := len(words) - count; i < len(words); i++ {
			words[i] = applySingleModifier(words[i], modifier)
		}

		// Rebuild text
		text = strings.Join(words, " ") + after
	}

	return text
}

//  Apply Single Modifier to One Word
func applySingleModifier(word, modifier string) string {

	switch modifier {

	case "hex":
		// Convert hex to decimal
		val, err := strconv.ParseInt(word, 16, 64)
		if err == nil {
			return strconv.FormatInt(val, 10)
		}

	case "bin":
		// Convert binary to decimal
		val, err := strconv.ParseInt(word, 2, 64)
		if err == nil {
			return strconv.FormatInt(val, 10)
		}

	case "up":
		return strings.ToUpper(word)

	case "low":
		return strings.ToLower(word)

	case "cap":
		return capitalize(word)
	}

	return word
}

//  Capitalize First Letter
func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}

	runes := []rune(word)
	runes[0] = unicode.ToUpper(runes[0])

	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}

	return string(runes)
}


//  Fix Punctuation Spacing
func fixPunctuation(text string) string {

	// Remove space before punctuation
	re := regexp.MustCompile(`\s+([.,!?:;])`)
	text = re.ReplaceAllString(text, "$1")

	// Ensure space after punctuation (if not multiple punctuation)
	re2 := regexp.MustCompile(`([.,!?:;])([^\s.,!?:;])`)
	text = re2.ReplaceAllString(text, "$1 $2")

	return text
}

//  Fix Quote Spacing
func fixQuotes(text string) string {

	// Remove space after opening quote
	re1 := regexp.MustCompile(`'\s+`)
	text = re1.ReplaceAllString(text, "'")

	// Remove space before closing quote
	re2 := regexp.MustCompile(`\s+'`)
	text = re2.ReplaceAllString(text, "'")

	return text
}

//  Fix "a" to "an" Before Vowel
func fixAAn(text string) string {

	words := strings.Fields(text)

	for i := 0; i < len(words)-1; i++ {
		if strings.ToLower(words[i]) == "a" {
			firstLetter := unicode.ToLower(rune(words[i+1][0]))
			if strings.ContainsRune("aeiou", firstLetter) {
				if words[i] == "A" {
					words[i] = "An"
				} else {
					words[i] = "an"
				}
			}
		}
	}

	return strings.Join(words, " ")
}