package main

func joinWithPunctuation(token []string) string {
	if len(token) == 0 {
		return ""
	}

	result := token[0]
	for i := 1; i < len(token); i++ {
		if isPunctuation(token[i]) {
			result += token[i]
		}
	}
	return result
}
