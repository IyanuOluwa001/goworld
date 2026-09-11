package main

import (
    // "errors"
    //"fmt"
    "strings"
)

// Question 1:// Write a function that repeats a string n times// repeatStr("ha", 3) -> "hahaha"// repeatStr("go", 0) -> ""
// func repeatStr(s string, n int) string {
//  return strings.Repeat(s, n)

// }

// // -------------------------------------------------------

// // Question 2:// Write a function that checks if a string is a palindrome (case-insensitive)// isPalindrome("Racecar") -> true// isPalindrome("Hello") -> false
// func isPalindrome(s string) bool {
//  word := strings.ToLower(s)
//  runes := []rune(word)

//  for i,j := 0, len(word)-1; i<j; i, j = i+1, j-1 {
//      if runes[i] == runes[j] {
//          return true
//      }
//  }
//  return false
// }

// // -------------------------------------------------------

// // Question 3:// Write a function that takes a slice of words and applies uppercase to the last N words// uppercaseLastN(["this", "is", "so", "exciting"], 2) -> ["this", "is", "SO", "EXCITING"]
// func uppercaseLastN(words []string, n int) []string {
//  for i := len(words)-n; i < len(words); i++ {
//      // return strings.Split(strings.ToUpper(words[i]), ",")
//      words[i] = strings.ToUpper(words[i])

//  }
//  return words

// }

// // -------------------------------------------------------

// // Question 4:// Write a function that returns the largest number in a slice// maxInSlice([3, 1, 9, 2]) -> 9// maxInSlice([]) -> error: slice is empty
// func maxInSlice(nums []int) (int, error) {

//  max := nums[0]
//  if len(nums) == 0 {
//      return 0, errors.New("slice is empty")
//  }
//  for _, num := range nums[1:] {
//      if num > max {
//          max = num
//      }

//  }
//  return max, nil

// }

// // -------------------------------------------------------

// // Question 5:// Write a function that counts how many times each word appears in a string// wordFreq("go is fun and go is easy") -> map[go:2 is:2 fun:1 and:1 easy:1]
// func wordFreq(s string) map[string]int {
//  words := strings.Fields(s)

//  count := make(map[string]int)
//  for _, word := range words {
//      count[word]++
//  }
//  return count

// }

// // -------------------------------------------------------

// // Question 6:// Write a function that removes all vowels from a string// removeVowels("hello world") -> "hll wrld"
// func removeVowels(s string) string {
//  var result strings.Builder
//  for _, ch := range s {
//      if !strings.ContainsRune("aeiou", ch) {
//          result.WriteString(string(ch))
//      }
//  }
//  return result.String()
// }

// // -------------------------------------------------------

// // Question 7:// Write a function that encrypts a string using a Caesar cipher// each letter is shifted by n positions in the alphabet// caesarCipher("abc", 1) -> "bcd"// caesarCipher("xyz", 2) -> "zab"
// func caesarCipher(s string, shift int) string {
//  var result strings.Builder
//  for _, ch := range s {
//      if ch >= 'a' && ch <= 'z' {
//          ch = ((ch-'a')-rune(shift)+26)%26 + 'a'
//          result.WriteString(string(ch))
//      } else if ch >= 'A' && ch <='Z' {
//          ch = ((ch-'A')-rune(shift)+26)%26 + 'A'
//          result.WriteString(string(ch))
//      }
//  }
//  return result.String()
// }

// func repeatStr(s string, n int) string {
//  // return strings.Repeat(s, n)
//  if n <= 0 {
//      return ""
//  }
//  return strings.Repeat(s, n)
// }

// // -------------------------------------------------------

// // Question 8:// Write a function that trims spaces inside single quotes// trimSpaces("' hello '") -> "'hello'"
func trimSpaces(s string) string {
    return strings.Join(strings.Fields(s), "")
}

// func isAnagram(word1 string, word2 string) bool {
//  word1 = strings.ToLower(word1)
//  word2 = strings.ToLower(word2)

//  if len(word1) != len(word2) {
//      return false
//  }

//  count := make(map[rune]int)

//  for _, ch := range word1 {
//      count[ch]++
//  }

//  for _, ch := range word2 {
//      count[ch]--
//      if count[ch] < 0 {
//          return false
//      }
//  }
//  return true
// }

// // -------------------------------------------------------

/*
func main() {
    // fmt.Println(repeatStr("ha", 3))
    // fmt.Println(isPalindrome("Racecar"))
    // fmt.Println(uppercaseLastN([]string{"this", "is", "so", "exciting"}, 2))
    // fmt.Println(removeVowels("hello world"))
    // fmt.Println(caesarCipher("Hello", 3))
    // fmt.Println(repeatStr("hello", 3))
    // fmt.Println(isAnagram("WORLD", "hello"))
    // fmt.Println(wordFreq("go is fun and go is easy"))
    // fmt.Println(maxInSlice([]int{3, 1, 9, 2}))
    fmt.Println(trimSpaces("' hello '"))

}
*/

