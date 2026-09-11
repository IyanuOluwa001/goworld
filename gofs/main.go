package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Run it this way: go run . \"text\" banner")
		return
	}

	input := strings.ReplaceAll(os.Args[1], "\\n", "\n")
	banner := os.Args[2]
	filename := banner + ".txt"

	lines, err := readBannerFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	
	asciiMap := buildAsciiMap(lines)
	renderAscii(input,asciiMap)
}


func readBannerFile(filename string) ([]string, error){
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
//Turns characters in my banner to a list
	lines := strings.Split(string(data), "\n")
	return lines, nil
}

func buildAsciiMap(lines []string) (map[rune][]string) {
	asciiMap := make(map[rune][]string)
	for ascii := 32; ascii <= 126; ascii++ {
		index := ascii - 32
		start := index * 9
		var block []string
		for i:= 0; i<8; i++{
			block = append(block, lines[start+i])
		}
		asciiMap[rune(ascii)] = block //ish
		//so that later we can do asciiMap['A']
	}
	return asciiMap
}

func renderAscii(input string, asciiMap map [rune][]string){
	lines := strings.Split(input, "\n")
	for _, line := range lines{
		if line == ""{
			fmt.Println()
			continue
		}
		for row := 0; row<8; row++{
			for _, char := range line{
				block, ok := asciiMap[char]
				if !ok{
				continue
			}
			fmt.Print(block[row])
		}
		fmt.Println()
		}
	}
}
