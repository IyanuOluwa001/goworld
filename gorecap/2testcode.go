package main

import (
	"fmt"
	"strconv"
)

func TestVar() {
	var username1 string = "Steve"
	var username2 string = "Jobs"
	var regno int = 20
	regno2 := 21

	fmt.Println(username1 + " and " + username2 +
		" have the reg no: " + strconv.Itoa(regno) +
		" and " + strconv.Itoa(regno2) + " respectively")
}
