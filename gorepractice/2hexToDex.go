package main

import "strconv"

func hexToDec(s string) (int64, error){
	val, err:= strconv.ParseInt(s, 16, 64)
	if err != nil{
		return 0, nil
	}
	return val, nil
}