package main

import (
	"ascii-art-justify/functions"
	"fmt"
	"os"
	"strings"
)

func main() {
	stringIndex := 1
	bannerIndex := 2
	var alignType string = "left"
	if len(os.Args) > 1 {
		if strings.HasPrefix(os.Args[1], "--align=") {
			parts := strings.Split(os.Args[1], "=")
			alignType = parts[1]
			stringIndex = 2
			bannerIndex = 3

			if alignType != "left" && alignType != "right" && alignType != "center" && alignType != "justify" {
				fmt.Println("Error: aligntype not handled")
				return
			}
		} else if strings.HasPrefix(os.Args[1], "-") {
			fmt.Println("Error: wrongly formed flag")
			return
		}
	}

	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("Example: go run . --align=right Hello standard")
		return
	}

	banner := "standard"
	text := ""

	if len(os.Args) > bannerIndex { //2 
		banner = os.Args[bannerIndex] //2
	}
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		fmt.Println("usage: go run . [string] [banner]")
		return
	}

	if len(os.Args) > stringIndex {
		text = os.Args[stringIndex]
	}

	if text == "" {
		return
	}
	if text == "\n" {
		fmt.Println()
		return
	}
	// alignType := "right"
	fmt.Print(functions.AsciiArt(text, alignType, banner))
}
