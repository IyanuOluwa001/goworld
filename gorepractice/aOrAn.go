package main

import "strings"

// 6.
func aOrAn1(s string) string {
	sIsLow := strings.ToLower(s)
	article := "a"
	for i, ch := range sIsLow {
		if i == 0 {
			vowels := []rune{'a', 'e', 'i', 'o', 'u', 'h'}
			for _, vowel := range vowels {
				if ch == vowel {
					article = "An"
					break
				}
			}
			break
		}
	}
	return article
}

//6b.
//Write a function that determines whether to use "a" or "an" before a given word // "apple" -> "an"
// "horse" -> "an" // "book" -> "a" // "honest" -> "an" (starts with h)

func aOrAn(nextWord string) string {
	if nextWord == "" {
		return "a"
	}

	word := strings.ToLower(nextWord)
	silentH := []string{"honest", "hour", "horse", "honor", "heir"}

	for _, w := range silentH {
		if strings.HasPrefix(word, w) {
			return "an"
		}
	}

	vowels := "aeiou"

	if strings.ContainsRune(vowels, rune(word[0])) {
		return "an"
	}

	return "a"
}

// 7.
func crudeAorAn(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words)-1; i++ {
		if (words[i] == "a" || words[i] == "A") && strings.ContainsRune("aeiouhAEIOUH", rune(words[i+1][0])) {
			words[i] += "n"
		}
	}
	return strings.Join(words, " ")
}
