package main

import (
	//"regexp"
	"strings"
)

// 9.
func removeSpaceBeforePunct(s string) string {
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' && i+1 < len(s) && strings.ContainsRune(".,!?:;", rune(s[i+1])) {
			continue
		}
		sb.WriteByte(s[i])
	}
	return sb.String()
}
