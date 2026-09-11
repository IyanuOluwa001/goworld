package main

import (
	//"regexp"
	"unicode"
)

// 8.
func fixPunctSpacing(test []string) string {
	strTest := ""
	for i, char := range test {
		isPunct := true
		for _, ch := range char {
			if !unicode.IsPunct(ch) {
				isPunct = false
			}
		}
		if isPunct || i == 0 {
			strTest += char
			continue
		}
		strTest += " " + char
	}
	return strTest
}
