package main

import "strings"

func stringReverser(s string) string {
	final := ""
	for _, ch := range s {
		final = string(ch) + final
	}
	return final
}

// 2b.
func reverse(s string) string {
	var sb strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteString(string(s[i]))
	}
	return sb.String()
}
