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

	if input == "" {
		return
	}

	result, err := asciiArt(input, banner)
	if err != nil {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		return
	}

	if outputfile != "" {
		err = os.WriteFile(outputfile, []byte(result), 0664)
		if err != nil {
			fmt.Println("Error Writing File")
			return
		} else {
			fmt.Println("Successful !!!")
		}

	} else {
		fmt.Print(result)
	}
}
