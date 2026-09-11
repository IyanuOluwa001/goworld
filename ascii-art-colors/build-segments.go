package main

import (
	"os"
	"strings"
)

// buildSegment() returns a slice of Segments data type that we declared in line 7 of main.go
func buildSegment(full, find string) []Segments { //What does this function do?

	//In the above code, segment is already defined, and so it is a data type.
	res := []Segments{} //we return a slice of segments.. (different parts and cuts)

	// the below checks, if the full string we are to get substring from is valid  (i.e .. --color=red "sub" "fullstring")
	if full == "" && (len(os.Args) >= 3) || isPresent(os.Args[len(os.Args)-1]) { //If full is absent yeah, we want to just say the find..is all true (args lack a substring)!
		res = append(res, Segments{text: find, colored: true}) // Add to the slice of segments and return true for all . you didnt provide a fullstring
		return res                                             //end the program here
	}

	//if the user provides the two strings.. then proceed to filter out
	if strings.Contains(full, find) { // if TRUE, code continues

		//getAll() returns all indexes of all submatch
		idx := getAll(full, find) //[0, 9 for example]
		prev := 0                 //prev is the variable we use to store the checkpoint

		for _, id := range idx { //loop through all starting value

			if prev < id { //checks for the strings before the target. if there arent any
				res = append(res, Segments{text: full[prev:id], colored: false}) //append all from the beginning to the char before the match, and make the value of colored FALSE
			}

			res = append(res, Segments{text: full[id : id+len(find)], colored: true}) //append the target index from start to (len(find)), find chars and make true

			prev = id + len(find) // update the value of prev so it can continue from there for the next match index
		}

		//if after add up prev is less than full, and after the loop from above
		if prev < len(full) {
			res = append(res, Segments{text: full[prev:], colored: false}) //append for the string after last colored string
		}
	}

	return res //return slices of segments to use to print the colors..
}

// getAll() returns all starting index of all matches
func getAll(full, find string) []int {
	offset := 0     //starting point, like "prev" to save checkpoint
	book := []int{} //what we return

	for { // an infinite loop
		i := strings.Index(full[offset:], find) //strings.Index is to give the found outcome index
		if i == -1 {                            //stops if condition TRUE, can only be -1 if no match if found anymore. so its possible through the loop
			break //end of the program
		}

		gap := offset + i        //checkpoint gap
		book = append(book, gap) //append to return value
		offset = gap + len(find)
	}

	return book //return value of book, which contains the first indexes of all matches
}

// a small function that checks if an argument is a banner
func isPresent(s string) bool {
	check := map[string]bool{
		"thinkertoy": true,
		"standard":   true,
		"shadow":     true,
	}

	for _, ok := check[s]; ok; {
		return true
	}
	return false
}
