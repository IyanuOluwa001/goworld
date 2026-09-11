package main

import (
	"strings"
	"unicode"
)

func Tokenize(text string) []string {
	var tokens []string
	var current strings.Builder
	inParentheses := false

	for _, r := range text {

		switch {

		case r == '(':
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}

			inParentheses = true
			current.WriteRune(r)

		case r == ')':

			current.WriteRune(r)
			tokens = append(tokens, current.String())

			current.Reset()
			inParentheses = false

		case unicode.IsSpace(r):

			if inParentheses {
				current.WriteRune(r)

			} else {

				if current.Len() > 0 {
					tokens = append(tokens, current.String())
					current.Reset()
				}
			}

		default:
			current.WriteRune(r)
		}
	}

	if current.Len() > 0 {
		tokens = append(tokens, current.String())
	}

	return tokens
}

func RebuildText(tokens []string) string {
	return strings.Join(tokens, " ")
}

func RemoveToken(tokens []string, index int) []string {
	return append(tokens[:index], tokens[index+1:]...)
}
