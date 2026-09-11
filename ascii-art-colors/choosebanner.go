package main

import (
	"fmt"
	"os"
	"strings"
)

// chooseBanner() function to choose the banner files
func chooseBanner() (string, error) {

	output := ""      //to assign the banner recognized
	def := "standard" // set the default file if valid and no file recognized
	inputs := os.Args

	switch len(inputs) { //using a switch statement to check. can also use an esle if statement
	case 3:
		if (inputs[2] != "thinkertoy" && inputs[2] != "standard" && inputs[2] != "shadow") && !(strings.HasPrefix(inputs[1], "--color=")) {
			break
		} else {
			output = def
		}
	case 4:
		// assign output to default if condition is met.
		if strings.Contains(inputs[3], inputs[2]) && (inputs[3] != "thinkertoy" && inputs[3] != "standard" && inputs[3] != "shadow") {
			output = def
			// assign output to the file passed
		} else if inputs[3] == "thinkertoy" || inputs[3] == "standard" || inputs[3] == "shadow" {
			output = inputs[3]
		} else {
			// else break out of the switch, return error
			break
		}
	case 5:
		if inputs[4] != "thinkertoy" && inputs[4] != "standard" && inputs[4] != "shadow" {
			break
		} else {
			output = inputs[4]
		}
	default:
		output = def
	}

	// if output is not empty state, as initialized return reading path
	if output != "" {
		store := fmt.Sprintf("./banner/%s.txt", output)
		return store, nil
	} else {
		// else if output wasnt assigned, meaning the args was invalid, return error
		return "", fmt.Errorf("ERROR: Invalid")
	}
}
