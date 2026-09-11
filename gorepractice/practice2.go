package main

import (
	"slices"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	out := [][]string{}
	for _, str1 := range strs {
		innerCont := make([]string, 0)
		innerCont = append(innerCont, str1)

		for _, str2 := range strs {
			if isPresentInSlice2d(out, str2) {
				continue
			}
			if isAnagram(str1, str2) && str1 != str2 {
				innerCont = append(innerCont, str2)
			}			
		}

		for _, word := range innerCont {
			if !isPresentInSlice2d(out, word) {
				out = append(out, innerCont)
				break
			}
		}
	}

	return out
}


func isPresentInSlice2d(sliceTwoD [][]string, str string) bool {
	for i := len(sliceTwoD); i > 0; i-- {
		if slices.Contains(sliceTwoD[i-1], str) {
			return true
		}
	}
	return false
}

func isAnagram(str1, str2 string) bool{
	for _, ch := range str1{
		if strings.ContainsRune(str2, ch){
			str2 = strings.Replace(str2, string(ch), "", 1)
			str1 = strings.Replace(str1, string(ch), "", 1)
		}
	}
	return str2 == str1
}