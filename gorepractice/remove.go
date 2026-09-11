package main

import (
	//"strconv"
)

func removeFSp(s string) string{
	for i, ch:= range s{
		if ch == ' '{
			return s[:1]+s[i+1:]
		}
	}
	return s
}



