package main

import (
	"fmt"
	"os"
	"strings"
)

/*
MAIN FUNCTION

This is where my program starts.

What my main function does:
1. Validate command-line arguments
2. Convert literal "\n" into actual newline characters
3. Load the banner file
4. Build the ASCII character map using the buildAsciiMap(lines)
5. Render the ASCII-art output; using renderAscii(input, asciiMap) 
*/
func main() {

	/*
		For this program, we expects exactly TWO arguments:
		os.Args[0] -> program name
		os.Args[1] -> input text
		os.Args[2] -> banner type (standard, shadow, thinkertoy)

		If the number of arguments is incorrect,
		we prints usage instructions and exits.
	*/
	if len(os.Args) != 3 {
		fmt.Println("Run it this way: go run . \"text\" banner")
		return
	}

	/*
		Handling the \n; because it is a literal, we have to convert it. E.g:
		go run . "Hello\nWorld" standard
		The shell passes this as the two characters "\" and "n".
		So we convert them into an actual newline character.
	*/
	input := strings.ReplaceAll(os.Args[1], "\\n", "\n")
	

	/*
		Banner style selected by the user.
		Example:
		standard
		shadow
		thinkertoy
	*/
	banner := os.Args[2]
	/*
		Construct the banner filename.
		Example:
		standard -> standard.txt
	*/
	filename := banner + ".txt"

	/*
		Read the banner file and return its contents
		as a slice of strings (each line of the file).
	*/
	lines, err := readBannerFile(filename)
	if err != nil {
		fmt.Println("Error: This file has nothing inside it", err)
		return
	}

	/*
		Convert the banner file into a map where:
		KEY   = character (rune)
		VALUE = slice of 8 ASCII-art lines

		This allows quick lookup when rendering characters.
	*/
	asciiMap := buildAsciiMap(lines)

	/*
		Render the ASCII-art representation of the input.
	*/
	renderAscii(input, asciiMap)
}

/*
readBannerFile

Reads the banner template file and returns all lines.

Example banner files:
standard.txt
shadow.txt
thinkertoy.txt

Each file contains ASCII-art definitions for
all printable ASCII characters.
*/
func readBannerFile(filename string) ([]string, error) {

	/*
		os.ReadFile reads the entire file into memory.

		It returns:
		- []byte containing the file content
		- an error if something went wrong
	*/
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	/*
		Convert the file content into a string
		and split it by newline characters.

		The result is a slice where each element
		represents one line from the file.
	*/
	//lines := strings.Split(string(data), "\n")
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")

	return lines, nil
}

/*
buildAsciiMap

Converts banner file lines into a map structure.

Map structure:
map[rune][]string

Example entry:
'A' to asciiArt A

Why this is useful:
It allows instant access to the ASCII-art
representation of any character.
*/
func buildAsciiMap(lines []string) map[rune][]string {

	/*
		Create an empty map.

		KEY   -> rune (character)
		VALUE -> slice of 8 ASCII-art rows
	*/
	asciiMap := make(map[rune][]string)

	/*
		Printable ASCII characters range from:

		ASCII 32 -> space
		ASCII 126 -> ~

		Total characters = 95
	*/
	for ascii := 32; ascii <= 126; ascii++ {

		/*
			Determine the index of the character
			inside the banner file.

			Example:
			'H' ASCII value = 72

			index = 72 - 32 = 40
		*/
		index := ascii - 32

		/*
			Each character occupies 9 lines in the file:

			8 lines of ASCII art
			1 empty separator line

			So the starting line for a character is:
		*/
		start := index * 9 + 1

		/*
			This slice will store the 8 ASCII rows
			for the current character.
		*/
		var block []string

		/*
			Extract the 8 ASCII-art lines
			and append them to the slice.
		*/
		for i := 0; i < 8; i++ {
			block = append(block, lines[start+i])
		}

		/*
			Store the character block in the map.

			rune(ascii) converts the ASCII number
			into the actual character.
		*/
		asciiMap[rune(ascii)] = block
	}

	return asciiMap
}

/*
renderAscii

This function prints the ASCII-art representation
of the input string.

Important concept:
ASCII-art characters are 8 lines tall.

Instead of printing character by character,
we print ROW BY ROW.
*/
func renderAscii(input string, asciiMap map[rune][]string) {

	/*
		Split the input string by newline characters.

		This allows handling multi-line input such as:

		Hello
		World
	*/
	lines := strings.Split(input, "\n")

	/*
		Process each line separately.
	*/
	for _, line := range lines {

		/*
			If the line is empty, print a blank line.
			This preserves multiple newline spacing.
		*/
		if line == "" {
			fmt.Println()
			continue
		}

		/*
			Each ASCII-art character has 8 rows.

			We iterate row-by-row.
		*/
		for row := 0; row < 8; row++ {

			/*
				Loop through every character in the line.
			*/
			for _, char := range line {

				/*
					Retrieve the ASCII-art block for the character.

					ok will be false if the character
					does not exist in the map.
				*/
				block, ok := asciiMap[char]
				if !ok {
					continue
				}

				/*
					Print the current row of the character.
				*/
				fmt.Print(block[row])
			}

			/*
				After finishing a row for all characters,
				move to the next line.
			*/
			fmt.Println()
		}
	}
}