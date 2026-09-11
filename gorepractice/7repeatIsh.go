package main

import (
	"strings"
)

func repeatStr(s string, n int) string{
	if n <= 0{
		return ""
	}
	return strings.Repeat(s,n)
}

//For repeat string
//declare the parameters of string and the number of times, then return string
//check if the number is less than zero or the count
//if it is, return empty string
//then return strings.Repeat(s,n)