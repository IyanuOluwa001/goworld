package functions

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func GetTerminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	res, err := cmd.Output()

	if err != nil {
		fmt.Println("hjhfjdb")
		os.Exit(1)
	}
	//
	parts := strings.Fields(string(res))

	termWidth, _ := strconv.Atoi(parts[1])

	return termWidth
}

func ApplyAlignment(lineArt, alignType string, wordLineArts []string, termWidth int) string {
	//Apply alignment
	if alignType == "left" || (alignType == "justify" && len(wordLineArts) == 1) {
		return lineArt
	}

	lineArtWidth := len(lineArt)
	diff := termWidth - lineArtWidth

	switch alignType {
	case "right":
		// apply right align
		lineArt = strings.Repeat(" ", diff) + lineArt
	case "center":
		// apply center
		lineArt = strings.Repeat(" ", diff/2) + lineArt
	case "justify":
		// apply justify
		var wordArtsWidth int
		for _, lineArt := range wordLineArts {
			wordArtsWidth += len(lineArt)
		}

		diff := termWidth - wordArtsWidth
		if diff <= 0 {
			return lineArt
		}

		numOfWordGaps := len(wordLineArts) - 1
		spacePerGap := diff / numOfWordGaps
		remSpace := diff % numOfWordGaps

		var justifiedLineArt string

		for i, wordArt := range wordLineArts {
			spaceToAdd := spacePerGap
			if remSpace > 0 {
				spaceToAdd++
				remSpace--
			}
			justifiedLineArt += wordArt
			if i != len(wordLineArts)-1 {
				justifiedLineArt += strings.Repeat(" ", spaceToAdd)
			}
		}

		return justifiedLineArt

	default:
		fmt.Println("align type not handled")
		return ""
	}

	return lineArt
}
