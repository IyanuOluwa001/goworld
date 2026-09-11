package main

import (
	"strings"
	//"unicode"
)

/*
func isPalindrome(s string) bool{
	left := 0
	right := len(s) - 1

	for left < right{
		if s[left] != s[right]{
			return false
		}
		left++
		right--
	}
	return true
}
*/

/*
func isPalindrome(s string) bool{
	var cleaned []rune
	for _, r:= range s{
		if unicode.IsLetter(r) || unicode.IsDigit(r){
			cleaned = append (cleaned, unicode.ToLower(r))
		}
	}
	left := 0
	right := len(cleaned) - 1

	for left < right{
		if cleaned[left] != cleaned[right]{
			return false
		}
		left++
		right--
	}
	return true
}
*/

func palindrome(s string) bool {
	s = strings.ToLower(s)
	oppStr := ""
	for _, ch := range s {
		oppStr = string(ch) + oppStr
	}
	return oppStr == s
}

/*
palindrome
- turn all text to lower
- set opptext to nothing
- set for statement to loop through all characters in the text
- reverse it
- and check if the reverse is same as the original
*/
