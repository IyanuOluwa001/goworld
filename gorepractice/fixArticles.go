package main

import (
	//"regexp"
	"strings"
	"unicode"
)

/*
func fixArticles(s string) string {
	re := regexp.MustCompile(`\b[Aa]\b\s+([aeiouAEIOU]\w*)`)
	return re.ReplaceAllString(s, "An $1")
}
*/

func fixArticles(s string) string {
	splitted := strings.Fields(s)
	for i, word := range splitted {
		word = strings.ToLower(word)
		if word == "an" && !strings.ContainsRune("aeiouuAEIOUH", rune(splitted[i+1][0])) {
			splitted[i] = string(splitted[i][0])
		} else if word == "a" && strings.ContainsRune("aeiouAEIOUH", rune(splitted[i+1][0])) {
			splitted[i] += "n"
		}
	}
	return strings.Join(splitted, " ")
}

//Write a function that processes a full sentence and fixes all "a" → "an" corrections
// "There it was. A amazing rock. A honest man. A book." -> "There it was. An amazing rock. An honest man. A book."

func isVowel(r byte) bool {
	switch unicode.ToLower(rune(r)) {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}

func stripPunct(word string) (string, string) {
	// separates word from trailing punctuation like "rock." -> ("rock", ".")

	i := len(word)
	for i > 0 && unicode.IsPunct(rune(word[i-1])) {
		i--
	}
	return word[:i], word[i:]
}

func isSilentH(word string) bool {
	silent := []string{"honest", "hour", "honor", "heir"}
	w := strings.ToLower(word)

	for _, s := range silent {
		if strings.HasPrefix(w, s) {
			return true
		}
	}
	return false
}

func fixArticles2(text string) string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return text
	}

	for i := 0; i < len(words)-1; i++ {
		if strings.EqualFold(words[i], "a") {
			nextWord, _ := stripPunct(words[i+1])

			if len(nextWord) == 0 {
				continue
			}
			first := nextWord[0]
			if isVowel(first) || isSilentH(nextWord) {
				// preserve original casing style
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
