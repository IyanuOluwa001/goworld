package main

type Modifier struct {
	Type  string // up, low, cap, hex, bin
	Count int    // default 1
}

func ProcessText(text string) (string, error) {
	
	tokens := Tokenize(text)

	tokens = ApplyTransformations(tokens)

	tokens = FixPunctuation(tokens)

	tokens = ApplyGrammarRules(tokens)

	return RebuildText(tokens), nil
}
