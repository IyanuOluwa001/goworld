package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Uppercaser(s string) {
	fmt.Println(strings.ToUpper(s))
}

func Lowercaser(s string) {
	fmt.Println(strings.ToLower(s))
}

func Hex2D(s string) (int64, error) {
	val, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return 0, nil
	}
	return val, nil
}

func Bin2D(s string) (int64, error) {
	val, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return 0, nil
	}
	return val, nil
}

func repeater(s string, n int) string {
	if n <= 0 {
		return ""
	}
	return strings.Repeat(s, n)
}

func WordCounter(sentence, word string) int {
	words := strings.Fields(sentence)
	count := 0

	for _, currentWord := range words {
		if strings.ToLower(currentWord) == strings.ToLower(word) {
			count++
		}
	}
	return count
}


