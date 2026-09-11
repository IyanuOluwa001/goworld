package main
import "strings"

func ApplyModifier(tokens []string, index int, mod Modifier) []string {

	// No previous words
	if index == 0 {
		return tokens
	}

	count := mod.Count
	if count <= 0 {
		count = 1
	}

	for j := 0; j < count; j++ {
		targetIndex := index - 1 - j

		if targetIndex < 0 {
			break
		}

		word := tokens[targetIndex]

		switch mod.Type {

		case "up":
			tokens[targetIndex] = strings.ToUpper(word)

		case "low":
			tokens[targetIndex] = strings.ToLower(word)

		case "cap":
			tokens[targetIndex] = Capitalize(word)

		case "hex":
			converted, err := HexToDecimal(word)
			if err == nil {
				tokens[targetIndex] = converted
			}

		case "bin":
			converted, err := BinToDecimal(word)
			if err == nil {
				tokens[targetIndex] = converted
			}
		}
	}

	return tokens
}