package main
import "unicode"

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
