#Exercise 1 — ASCII Understanding

Write a Go program that prints the ASCII value of a character.

Example:

Input: A
Output: 65

Hint:

int('A')

#Exercise 2 — Map Practice

Create a map like this:

map[rune]string

Example output:

A -> Apple
B -> Banana
C -> Cat

#Exercise 3 — Slice Practice

Write a program that prints this:

Hello
World

From:

[]string{
"Hello",
"World",
}

#Exercise 4 — Nested Loops

Print this pattern:

11111
22222
33333
44444
55555

Using nested loops.

This prepares you for ASCII rendering logic.

Exercise 5 (Very Important)

Write a small Go program:

1 read standard.txt
2 split it into lines
3 print the number of lines

Expected result:

855 lines

Why?

95 characters × 9 lines = 855

Exercise 6

Print the first 10 lines of the banner file.

This helps you visualize the data structure.

Exercise 7

Write a function:

getCharBlock(lines []string, char rune)

Return:

[]string (8 lines of ASCII art)

Example:

getCharBlock(lines, 'A')

returns the 8 ASCII rows of A.

9. Exercise 1

Write a program that:

1 reads standard.txt
2 builds asciiMap
3 prints asciiMap['B']

Goal:

Understand map access.

10. Exercise 2

Print the ASCII art of every letter from A to Z.

Pseudo:

for char := 'A'; char <= 'Z'; char++ {
    print asciiMap[char]
}

This helps you understand the map structure deeply.

11. Exercise 3 (Important)

Write a function:

func getAsciiBlock(asciiMap map[rune][]string, char rune) []string

Return the ASCII representation of a character.

Example:

getAsciiBlock(asciiMap, 'H')

15. Exercise 1 (Very Important)

Write a small program that renders:

ABC

using your asciiMap.

Goal: understand row-based printing.

16. Exercise 2

Modify renderAscii() so that it prints:

Hello\nWorld

correctly.

Expected output:

HELLO ASCII ART

WORLD ASCII ART

17. Exercise 3 (Debug Skill)

Add debug prints:

fmt.Println("Rendering row:", row)

to observe how the renderer works.

Exercise 1

Modify the program to support:

go run . "Hello" shadow

and print shadow-style ASCII art.

Exercise 2

Write a function:

func isValidAscii(char rune) bool

Return true only for characters between 32–126.

Exercise 3

Write a test that checks:

asciiMap['Z']

returns exactly 8 lines.

Exercise 4 (Advanced)

Modify the renderer so that it returns a string instead of printing directly.

Example:

func renderAscii(input string, asciiMap map[rune][]string) string

This makes testing much easier.


## I don't understand this yet:
We convert the file into a slice:

[]string

Example:

lines := strings.Split(string(data), "\n")

Now we can access:

lines[0]
lines[1]
lines[2]
...

This makes indexing easy.

## to get 
func buildAsciiMap(lines []string) map[rune][]string {
	asciiMap := make(map[rune][]string)

	for ascii := 32; ascii <= 126; ascii++ {

		index := ascii - 32
		start := index * 9

		var block []string

		for i := 0; i < 8; i++ {
			block = append(block, lines[start+i])
		}

		asciiMap[rune(ascii)] = block
	}

	return asciiMap
}

## to get 2
func renderAscii(input string, asciiMap map[rune][]string) {

	for row := 0; row < 8; row++ {

		for _, char := range input {

			block, ok := asciiMap[char]
			if !ok {
				continue
			}

			fmt.Print(block[row])
		}

		fmt.Println()
	}
}

