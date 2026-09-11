package main

import "strings"

func capitalizer(word string) string {
	final := ""
	for i, ch := range word {
		if i == 0 {
			final += strings.ToUpper(string(ch))
			continue
		}
		final += string(ch)
	}
	return final
}

func capitaliZer(s string) string {
	if len(s) == 0 {
		return s
	}

	runes := []rune(s)
	for i := 1; i < len(runes); i++ {

		if runes[0] >= 'a' && runes[0] <= 'z' {
			runes[0] -= 32
		}
	}
	return string(runes)
}
