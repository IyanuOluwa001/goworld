package main

import (
	"fmt"
	"os"
	"strings"
)

func auth() ([]string, bool, error) { //

	if asciiArt() { // ascii art returns a bool. TRUE is args is ascii-art, FALSE otherwise
		return nil, true, nil
	}

	path, err := chooseBanner() // a function that checks the banner file to use, and returns the path with the banner included and an error
	if err != nil {             //if there is an error
		// fmt.Println(err)
		return nil, false, err
	}

	filedata, err := os.ReadFile(path)            //error if path isnt available or invalid
	data := strings.Split(string(filedata), "\n") // if path is valid, extract data and split by new line, into slices
	// data1 := strings.ReplaceAll(string(filedata), "\n")

	//can be removed
	if err != nil {
		// fmt.Println("ERROR: Invalid file to read")
		return nil, false, fmt.Errorf("ERROR: Invalid file to read")
	}

	return data, false, nil // if the code gets to here; the arguments isnt ascii art
}

// checkos()  checks and returns 3 strings and an error is presnt
func checkos() (string, string, string, error) {
	input := os.Args //input will be declared to be used as os.Args

	if len(input) < 2 || len(input) > 5 { //guard for check
		return "", "", "", fmt.Errorf("ERROR: Number of arguments passed is invalid")
	}

	color := input[1]
	find := input[2]
	full := "" // full is initialized empty because we are not sure if its present or not

	// this block of code assigns full a value if the valid conditions are met
	if len(input) >= 4 && input[3] != "" && len(input[2]) > 0 && strings.Contains(input[3], find) {
		full = input[3]

	} else if len(input) < 2 {
		return "", "", "", fmt.Errorf("ERROR: Incomplete number of argument")
	} else if len(input) == 4 && !(strings.Contains(input[3], find)) && !isPresent(input[3]) {
		return "", "", "", fmt.Errorf("ERROR: Cannot find substring in main")
	}

	// constants of the colors that can be used
	colors := map[string]string{
		"red":   "\033[31m",
		"green": "\033[32m",
		"blue":  "\033[34m",
	}

	//check if color code is valid
	if strings.HasPrefix(color, "--color=") && len(color) > 8 {

		//divide to strings if true and take the index of 1 (second part)
		w := strings.Split(color, "=")
		color = w[1] //assign it to color

		if val, ok := colors[color]; ok { //if the color passed in is found in the colors map
			color = val //assign finally back
		} else {
			// color is not valid, return an error back
			return "", "", "", fmt.Errorf("ERROR: Invalid color code passed in")
		}
	} else {

		//if color index is not valid, return and stop the program
		return "", "", "", fmt.Errorf("ERROR: Invalid color code passed in")
	}

	//if fullstring is present and substring can be found in mainstring
	if full != "" && !strings.Contains(full, find) { //if nto present
		return "", "", "", fmt.Errorf("ERROR: Substring cannot be found in string")
	}

	// returns assigned values if all is present..
	return color, find, full, nil
}
