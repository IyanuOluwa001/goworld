package main

func buildAsciiMap(lines []string) map[rune][]string {
	asciiMap := make(map[rune][]string)
	for ascii := 32; ascii <= 126; ascii++ {
		index := ascii - 32
		start := index*9 + 1
		var block []string
		for i := 0; i < 8; i++ {
			block = append(block, lines[start+i])
		}
		asciiMap[rune(ascii)] = block //ish
		//so that later we can do asciiMap['A']
	}
	return asciiMap
}
