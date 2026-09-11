package main


func ApplyTransformations(tokens []string) []string {
	for i := 0; i < len(tokens); i++ {

		if IsModifier(tokens[i]) {

			mod := ParseModifier(tokens[i])

			tokens = ApplyModifier(tokens, i, mod)

			tokens = RemoveToken(tokens, i)

			i--
		}
	}
	return tokens
}
