/*
ToUpper
ToLower
palindrome
removespace
bintodec
hexToDec
repeatstrg
capitalizer of final+string(ch)
countword of words:= strings.Fields(s)
*/

package main

import "strings"

func capiT(s string) string{
	final:= ""
	for i, ch:= range s{
		if i==0{
		final+=strings.ToUpper(string(ch))
		continue
	}
	final+=string(ch)
}
return final
}

//1.
//so use a for loop with a single comma, use ch
//to range through s
//check if it is ch is empty, if it is remove
//return

//2.
//for hextodec, we accept string and output int64
//and error. we declare val and error, do the
//conversation.
//check err and return 0, then return val
