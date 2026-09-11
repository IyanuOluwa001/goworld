package main

import (
	//"regexp"
	"strings"
)

// if you want to try lowering while counting nth times from front,
// just uncomment the lines of code i commented inside the function
// and then comment the ones untop them. For now it's is the normal
// lowering from the back
func nthWordToLowerFront(words []string, n int) string {
	switch {
	case n < 0:
		n = 0
	case n > len(words):
		n = len(words)
	}
	// toBeLower := words[:n]
	toBeLower := words[len(words)-n:]
	startIndex := len(words) - n
	for i, ch := range toBeLower {
		words[startIndex+i] = strings.ToLower(ch)
		// words[i] = strings.ToLower(ch)
	}
	return strings.Join(words, " ")
}
