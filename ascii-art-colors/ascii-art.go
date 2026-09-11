package main

import (
	"fmt"
	"os"
	"strings"
)

// ascii art code but returns a boolean TRUE OR FALSE
func asciiArt() bool {

	j := os.Args
	// out := false
	if (len(j) > 3) || (len(j) == 3 && strings.HasPrefix(j[1], "--color=")) || (len(j) == 3 && !isPresent(j[2])) {
		// out = true
		return false
	}
	data, err := readFile()
	if err != nil {
		fmt.Println(err)
		// return
	}

	input := os.Args[1]

	file := strings.Split(string(data), "\n") //splits the data by every new line --> 655

	userinput := strings.ReplaceAll(input, "\\n", "\n") //returns a string back

	maininput := strings.Split(userinput, "\n")

	for _, word := range maininput {
		if word == " " {
			fmt.Println()
			continue
		}

		for row := 0; row <= 8; row++ {
			for _, ch := range word {
				start := int(ch-32) * 9
				fmt.Print(file[start+row]) //because row iterates and adds plus one in every loop.
			}
			fmt.Println()
		}
	}

	return true
}

// function to read and print ascii-art default
func readFile() ([]byte, error) {
	default_val := "standard"
	lent := len(os.Args)
	next := ""

	if lent < 2 {
		return nil, fmt.Errorf("ERROR: Incomplete argument")
	}

	switch lent {
	case 3:
		option := os.Args[2]

		if option != "standard" && option != "thinkertoy" && option != "shadow" {
			return nil, fmt.Errorf("ERROR: No file exist as such names")
		}
		next = option
	case 2:
		next = default_val
	default:
		return nil, fmt.Errorf("ERROR: Not required arguments")
	}

	read, err := os.ReadFile(fmt.Sprintf("./banner/%s.txt", next))
	return read, err
}
