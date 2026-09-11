package main

func ProcessText(text string) (string, error) {

	tokens := Tokenize(text)

	tokens = ApplyTransformations(tokens) //in here

	tokens = FixPunctuation(tokens)

	tokens = ApplyGrammarRules(tokens) //in here

	return RebuildText(tokens), nil
}

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

func ApplyGrammarRules(tokens []string) []string {
	for i := 0; i < len(tokens)-1; i++ {

		if IsArticleA(tokens[i]) && StartsWithVowelOrH(tokens[i+1]) {
			tokens[i] = FixArticleCase(tokens[i])
		}
	}
	return tokens
}
