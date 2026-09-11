package main

func FixPunctuation(tokens []string) []string {

	// Convert tokens to string
	text := RebuildText(tokens)

	// Apply string-based fixes
	text = FixStandardPunctuation(text)
	text = FixApostrophes(text)

	// Convert back to tokens
	return Tokenize(text)
}