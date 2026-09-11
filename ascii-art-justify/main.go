package main

import (
	"ascii-art-justify/asciifunctions"
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

			if alignType != "center" && alignType != "right" && alignType != "justify" {
				fmt.Println("wrong aligntype")
				return
			}
		} else if strings.HasPrefix(os.Args[1], "-") {
			fmt.Println("wrong flag")
		}
	}

	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("Example: go run . --align=right something standard")
	}

	text := ""
	if len(os.Args) > stringIndex {
		text = os.Args[stringIndex]
	}

	banner := "standard"

	if banner != "standard" && banner != "thinkertoy" && banner != "shadow" {
		fmt.Println("wrong banner")
		return
	}
	if len(os.Args) > bannerIndex {
		banner = os.Args[bannerIndex]
	}

	if text == "" {
		return
	}

	if text == "\n" {
		fmt.Println()
	}

	fmt.Print(asciifunctions.AsciiArt(text, alignType, banner))
}
