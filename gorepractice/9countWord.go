package main

import (
	"strings"
)

func countWord(s, target string) int{
	words:= strings.Fields(s)
	count := 0
	for _, word:= range words{
		if strings.ToLower(word)==strings.ToLower(target){
		//		if strings.EqualFold(word, target){
			count++
		}
	}
	return count
}

func countWords(sentence, word string) int{
	allLower:= strings.ToLower(sentence)
	allWords:= strings.Fields(allLower)
	count := 0
	for _, this:= range allWords{
		if this==word{
			count++
		}
	}
	return count
}

