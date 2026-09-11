package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		return
	}

	var input, banner, flag, outputfile string

	banner = "standard"

	if len(os.Args) == 2 {
		input = os.Args[1]
	}

	if len(os.Args) == 3 {
		input = os.Args[1]
		banner = os.Args[2]
	}

	if len(os.Args) == 4 {
		flag = os.Args[1]
		input = os.Args[2]
		banner = os.Args[3]

	if strings.HasPrefix(flag, "--output=") {
		outputfile = strings.TrimPrefix(flag, "--output=")
	} else {
		fmt.Println("Invalid flag usage")
		return
	}
	}


	data, err := os.ReadFile(banner + ".txt")
	if err != nil {
		fmt.Println("Error Read File")
		return
	}

	text := strings.Split(string(data), "\n")
	word := strings.Split(input, "\\n")

	var result strings.Builder

	for _, words := range word {

		for i := 0; i < 8; i++ {
			for _, character := range words {
				index := (int(character)-32)*9 + i
				result.WriteString(text[index])
			}
			result.WriteString("\n")
		}
	}

	if outputfile != "" {
		err = os.WriteFile(outputfile, []byte(result.String()), 0664)
		if err != nil {
			fmt.Println("Error Writen File")
			return
		} else {
			fmt.Println("Successful !!!")
		}

	} else {
		fmt.Print(result.String())
	}

}
