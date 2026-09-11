package main

import (
	"fmt"
)

func ageDilemma() {
	// 1. Print the prompt first. 
	// Using fmt.Print instead of Println keeps the blinking cursor on the same line.
	fmt.Print("Insert a time of the day: ")

	// 2. Create an integer variable to hold the user's answer
	var time int

	// 3. Tell Go to pause and wait for the user to type something.
	// The '&' symbol tells Scanln exactly where in memory to save the number.
	_, err := fmt.Scanln(&time)
	if err != nil {
		fmt.Println("Error: That wasn't a valid number!")
		return
	}

	// 4. Run your switch logic on the number they typed
	switch {
	case time < 12:
		fmt.Println("Good morning!")
	case time < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good evening")
	}
}