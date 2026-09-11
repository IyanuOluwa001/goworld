package main

import "unicode"

func isPunctuation(s string) bool {
	var isPunct bool
	for _, ch := range s {
		isPunct = unicode.IsPunct(ch)
		if isPunct {
			break
		}
	}
	return isPunct
}

func isPunctuation2(s string) bool {
	return s == "," || s == "." || s == "!" || s == "?" || s == ";"
}

// 14b.
func checkPunc(s string) bool {
	if len(s) != 1 {
		return false
	}

	punctuations := map[rune]bool{
		',':  true,
		'.':  true,
		'!':  true,
		'?':  true,
		';':  true,
		':':  true,
		'-':  true,
		'_':  true,
		'\'': true,
		'"':  true,
	}

	r := rune(s[0])
	return punctuations[r]
}
