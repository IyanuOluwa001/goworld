package main

func ApplyGrammarRules(tokens []string) []string {
	for i := 0; i < len(tokens)-1; i++ {

		if IsArticleA(tokens[i]) && StartsWithVowelOrH(tokens[i+1]) {
			tokens[i] = FixArticleCase(tokens[i])
		}
	}
	return tokens
}
