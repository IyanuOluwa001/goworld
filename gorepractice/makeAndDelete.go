package main

import "fmt"

func MapTester() {
	config := map[string]int{
		"connections":     100,
		"timeout_seconds": 30,
		"retry":           3,
	}

	//for key, value := range config {
	for key := range config {
		//fmt.Printf("Config key '%s' has value %d\n", key, value)
		fmt.Println("Active settings:", key)
	}
}
