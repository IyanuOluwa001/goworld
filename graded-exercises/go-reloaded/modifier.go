package main

import (
	"strconv"
	"strings"
	"unicode"
)

type Modifier struct {
	Type  string
	Count int
}

func IsModifier(token string) bool {

	if len(token) < 3 {
		return false
	}

	if token[0] != '(' || token[len(token)-1] != ')' {
		return false
	}

	return true
}

func ParseModifier(token string) Modifier {

	content := strings.TrimPrefix(token, "(")
	content = strings.TrimSuffix(content, ")")

	parts := strings.Split(content, ",")

	mod := Modifier{
		Type:  strings.TrimSpace(parts[0]),
		Count: 1,
	}

	if len(parts) > 1 {

		count, err := strconv.Atoi(strings.TrimSpace(parts[1]))

		if err == nil && count > 0 {
			mod.Count = count
		}
	}

	return mod
}

func ApplyModifier(tokens []string, index int, mod Modifier) []string {

	if index == 0 {
		return tokens
	}

	count := mod.Count

	for j := 0; j < count; j++ {

		target := index - 1 - j

		if target < 0 {
			break
		}

		word := tokens[target]

		switch mod.Type {

		case "up":
			tokens[target] = strings.ToUpper(word)

		case "low":
			tokens[target] = strings.ToLower(word)

		case "cap":
			tokens[target] = Capitalize(word)

		case "hex":

			if v, err := HexToDecimal(word); err == nil {
				tokens[target] = v
			}

		case "bin":

			if v, err := BinToDecimal(word); err == nil {
				tokens[target] = v
			}
		}
	}

	return tokens
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
