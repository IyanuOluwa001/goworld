package main

import (
	"fmt"
)

var (
	a int
	b int = 2
	c     = 3
)

func NewTest() {
	a = 1
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
