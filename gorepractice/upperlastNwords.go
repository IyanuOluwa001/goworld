package main

// 13.
func uppercaseLastNwords(s []string, n int) []string {
	if n <= 0 {
		return s
	}

	start := len(s) - n
	if start < 0 {
		start = 0
	}

	for i := start; i < len(s); i++ {
		runes := []rune(s[i])
		for j := 0; j < len(runes); j++ {
			if runes[j] >= 'a' && runes[j] <= 'z' {
				runes[j] -= 32
			}
		}
		s[i] = string(runes)
	}
	return s
}
