package main

import (
	"strconv"
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
			// flush current word
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

func HexToDecimal(s string) (string, error) {
	n, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(n, 10), nil
}

func BinToDecimal(s string) (string, error) {
	n, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(n, 10), nil
}

func Capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	runes := []rune(word)
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}
