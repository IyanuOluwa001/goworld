package main

import (
	"strings"
)

func removeVowels(s string) string {
	var b strings.Builder
	for _, r := range s {
		if !strings.ContainsRune("aeiouAEIOU", r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}