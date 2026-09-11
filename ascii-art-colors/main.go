package main

import (
	"fmt"
)

type Segments struct { //created a new type that stores a string and a bool, to filter colored and uncolored
	text    string
	colored bool
}

func main() { //the start of the main file

	//auth() function returns readfile data in strings, check for asciiart option, returns false if not an ascii arguments. err for any return error
	data, check, err := auth()

	if err != nil { // if error is not nil, then error is true
		fmt.Println(err) //stops the program and returns matching error
		return           //instances of return stops the program
	}

	if check { //if true : return (stops), the argument passed was an ascii-art value
		return
	}

	// checkos() returns back 3 strings and 1 error
	color, find, full, err := checkos() //matches the var
	if err != nil {
		fmt.Println(err)
		return
	}

	// buildString() returns a segment data type as declared at the top. Segment is a "string" and "bool" type.
	seg := buildSegment(full, find)

	if len(seg) == 0 { //if there is no match or invalid
		fmt.Println("ERROR: Substring and string dosent match") //print this
		return
	}

	fmt.Println(seg)   // for de-bugging purposes
	reset := "\033[0m" // reset is the closing tag of an opening color

	for row := 1; row <= 8; row++ { // just like ascii-art, print out uses 8 lines (rows) for each character

		for _, id := range seg { // segment is a {string : bool} type, loops through the slice
			if id.text == "\n" { //this is meant to check for new lines in the input passed in by the user
				if row == 1 {
					fmt.Println()
				}
				continue //continues the code..
			}

			for _, ch := range id.text { //each index values of the full seg slices we created earlier
				start := int(ch-32) * 9 // this gets the first line of each character
				if id.colored {         // since the segment data type is {string : bool}, check if colored is true ..
					fmt.Print(color + data[start+row] + reset) // add "color" opening to the section and close "reset" to close the tag so it dosent spread to the next char
				} else {
					fmt.Print(data[start+row]) //else print default
				}
			}
		}
		fmt.Println() // moves to the next line until row is = 8
	}
}
