package main

import (
	"fmt"
)

func main() {
	//fmt.Println(binToDec("1111"))
	//fmt.Println(hexToDec("11a"))
	//fmt.Println(stringReverser("richarD backwardS almosT lookS likE draculA"))

	//fmt.Println(isPunctuation(":"))
	//fmt.Println(isPunctuation("ball.")) //also works for hasPunctuation

	//test1 := []string{"i", "am", "MICHEAL", "and", "i", "AM", "A", "BOY"}
	//fmt.Println(nthWordToLowerFront(test1, 3))

	//fmt.Println(capitalizer("i 😋lord"))
	//fmt.Println(aOrAn("Don't quit"))

	//test2 := []string{"Hello", ",", "world", "!", "This", "is", "Micheal", "."}
	//fmt.Println(fixPunctSpacing(test2))

	//fmt.Println(fixSingleQuote("' Subarashi       '"))
	//fmt.Println(fixSingleQuote("I ' am ' a winners"))

	//fmt.Println(fixArticles("She is a engineer and a doctor"))
	//fmt.Println(fixArticles("A hour rock"))

	//fmt.Println(removeFirstSpace("helloworld"))

	//fmt.Println(palindrome("madam"))
	//1. After today
	//fmt.Println(isPalindrome("madam"))
	//fmt.Println(isPalindrome("hello"))

	//2.
	//fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
	//fmt.Println(isPalindrome("No lemon, no melon"))             // true
	//fmt.Println(isPalindrome("hello"))

	//fmt.Println(reverse("Richard"))

	//fmt.Println(repeatStr("ha", 3))
	//Uppercaser("Hello")
	//Lowercaser("MOTO")
	//MapTester()

	/*
	//isPunctuation
	tests := []string{",", "!", "x"}
	for _, t := range tests {
		fmt.Printf("q->%v\n", t, isPunctuation(t))
	}

	tokens := []string{"hello", ",", "world", "!"}
	fmt.Println(joinWithPunctuation(tokens)) // hello, world!

	test2 := []string{"I", "am", "angry", "!"}
	fmt.Println(joinWithPunctuation(test2)) // I am angry!

	test3 := []string{"I", "just", "needed", "to", "remember", "slice", "syntax", "that", "is", "why", "?"}
	fmt.Println(joinWithPunctuation(test3))
*/
/*
	// Explain what strings.Fields() does and why it's useful in this project.
	fmt.Println("strings.Fields() is  package that helps you to save and separate substrings in a string by")
	fmt.Println("separating them with whitespaces... \nIt is important when you want to work on individual")
	fmt.Println("substrings, in a sense, mainupulating your way through:")
*/
/*
	// Another
	token := []string{"hello", ",", "world", "!"}
	fmt.Println(joinWithPunctuation(token)) // hello, world!

	test4 := []string{"I", "am", "angry", "!"}
	fmt.Println(joinWithPunctuation(test4)) // I am angry!

	test5 := []string{"I", "just", "needed", "to", "remember", "slice", "syntax", "that", "is", "why", "?"}
	fmt.Println(joinWithPunctuation(test5))
*/
/*
	// Another
	words := []string{"apple", "horse", "book", "honest"}

	for _, w := range words {
		fmt.Printf("%s %s\n", aOrAn(w), w)
	}
*/

	// Another
	//fmt.Println("\nOutput:")
/*
	// Another
	test6 := []string{
		"' awesome '",
		"' hello world '",
		"nothing here",
		"mix ' spaced text ' outside",
		"'multiple   spaces   here'",
	}

	for _, t := range test6 {
		fmt.Println(fixSingleQuotes(t))
	}
}
*/

fmt.Println(countWord("He is a bad he goat", "he"))
//14.
/*
func checkPunct (s string) bool{
	for i, val:= range s{
		if strings.ContainsRune(".,;:!?"){
			return true
		}
	}
}
*/

//Todo:
//check my goreload to see the functions I used
//check the doc i sent to israel on goreload recoding
//What is base and bitsize, why use strvonc.ParseInt
}

/*
	test4 := []string{"eat", "tea", "wept", "tan", "dog", "ate", "wet", "cry", "nat", "bat", "cat"}
	fmt.Println(groupAnagrams(test4))
	fmt.Println(isAnagram("cat", "act"))
*/
