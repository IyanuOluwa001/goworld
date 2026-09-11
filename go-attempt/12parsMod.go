package main
import (
	"strconv"
	"strings"
)

func ParseModifier(token string) Modifier {

	// Remove surrounding parentheses
	content := strings.TrimPrefix(token, "(")
	content = strings.TrimSuffix(content, ")")

	// Split by comma (if exists)
	parts := strings.Split(content, ",")

	mod := Modifier{
		Type:  "",
		Count: 1, // default
	}

	// First part is modifier type
	mod.Type = strings.TrimSpace(parts[0])

	// If count exists
	if len(parts) > 1 {
		countStr := strings.TrimSpace(parts[1])
		count, err := strconv.Atoi(countStr)
		if err == nil && count > 0 {
			mod.Count = count
		}
	}

	return mod
}