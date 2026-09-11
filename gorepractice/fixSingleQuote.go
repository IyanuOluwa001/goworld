package main

import "strings"

func fixSingleQuote(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' && i+1 < len(s) && s[i+1] == '\'' {
			continue
		}

		if s[i] == '\'' && i+1 < len(s) && s[i+1] == ' ' {
			result += "'"
			i++
			continue
		}

		result += string(s[i])
	}

	return result
}

/*
package main

import "fmt"

func fixSingleQuote(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		// skip space before '
		if s[i] == ' ' && i+1 < len(s) && s[i+1] == '\'' {
			continue
		}

		// skip space after '
		if s[i] == '\'' && i+1 < len(s) && s[i+1] == ' ' {
			result += "'"
			i++ // skip the space after '
			continue
		}

		result += string(s[i])
	}

	return result
}

func main() {
	fmt.Println(fixSingleQuote("I ' am ' a winners"))
}

*/

/*
func fixSingleQuote(s string) string {
	splittedTest := strings.Fields(s)
	final := ""
	for i, char := range splittedTest {
		if char == " " && (splittedTest[i-1] == "'" || splittedTest[i+1] == "'") {
			continue
		}
		final += char
	}
	return final
}
*/

/*
package main

import "strings"

func fixSingleQuote(s string) string {
	splittedTest := strings.Fields(s)
	var final strings.Builder
	for i, char := range splittedTest {
		if char == " " && (splittedTest[i-1] == "'" || splittedTest[i+1] == "'") {
			continue
		}
		final.WriteString(char)
	}
	return final.String()
}

*/
// Write a function that fixes spacing inside single quotes // "' awesome '" -> "'awesome'" // "' hello world '" -> "'hello world'"
// I have deleted the import, find another code to replace this solution
func fixSingleQuotes(text string) string {
	var result strings.Builder
	inQuotes := false
	var buffer strings.Builder

	for _, r := range text {
		if r == '\'' {
			if inQuotes {
				// closing quote → process buffer
				content := strings.TrimSpace(buffer.String())
				result.WriteRune('\'')
				result.WriteString(content)
				result.WriteRune('\'')
				buffer.Reset()
				inQuotes = false
			} else {
				// opening quote
				inQuotes = true
				buffer.Reset()
			}
			continue
		}

		if inQuotes {
			buffer.WriteRune(r)
		} else {
			result.WriteRune(r)
		}
	}

	// in case of unmatched quote, flush safely
	if buffer.Len() > 0 {
		result.WriteString(buffer.String())
	}

	return result.String()
}

func fixSingleQuotes2(s string) string {
	s = strings.ReplaceAll(s, "' ", "'")
	s = strings.ReplaceAll(s, " '", "'")
	return s
}
