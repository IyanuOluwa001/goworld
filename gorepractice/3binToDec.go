package main

import "strconv"

func binToDec(s string) (int64, error){
	val, err:= strconv.ParseInt(s, 2, 64)
	if err!=nil{
		return 0, nil
	}
	return val, nil
}