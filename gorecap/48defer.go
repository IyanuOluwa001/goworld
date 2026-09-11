package main

import "fmt"

func deferer() {
	fmt.Println("counting...")

	for i:=0; i<10; i++{
		defer fmt.Println(i)
	}

	defer fmt.Println("done")
}