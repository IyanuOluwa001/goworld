package main

import (
	"strings"
)

func ceaserCipher(s string, n int) string {
	var b strings.Builder
	for _, ch := range s {
		switch {
		case ch >= 'a' && ch <= 'z':
			b.WriteRune('a' + (ch-'a'+rune(n))%26)
		case ch >= 'A' && ch <= 'Z':
			b.WriteRune('A' + (ch-'A'+rune(n))%26)
		default:
			b.WriteRune(ch)
		}
	}
	return b.String()
}
