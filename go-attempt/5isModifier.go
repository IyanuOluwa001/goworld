package main

func IsModifier(token string) bool {
	if len(token) < 3 {
		return false
	}

	if token[0] != '(' || token[len(token)-1] != ')' {
		return false
	}

	return true
}

/*
func IsModifier(word string) bool {
	return len(word) > 2 && word[0] == '(' && word[len(word)-1] == ')'
}
*/