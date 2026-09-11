package functions

import (
	"fmt"
	"os"
	"strings"
)

func AsciiArt(text, alignType, banner string) string {
	data, err := os.ReadFile(banner + ".txt")
	if err != nil {
		fmt.Println("error reading file:", err)
		return ""
	}
	bannerdata := strings.Split(string(data), "\n")

	var res string
	termWidth := GetTerminalWidth()

	for i := 0; i < 8; i++ {
		var lineArt string
		var wordLineArt string
		var wordLineArt2 []string

		for _, char := range text {
			asciidx := int(char-' ')*9 + 1 + i
			charidx := bannerdata[asciidx]

			lineArt += charidx

			if char == ' ' {
				wordLineArt2 = append(wordLineArt2, wordLineArt)
				wordLineArt = ""
				continue
			}

			wordLineArt += charidx
		}
		wordLineArt2 = append(wordLineArt2, wordLineArt)

		// fmt.Println(wordLineArt2, len(wordLineArt2))
		lineArt = ApplyAlignment(lineArt, alignType, wordLineArt2, termWidth)
		res += lineArt + "\n"
	}
	return res
}
