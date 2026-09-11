package main

import (
	"fmt"
	"math"
	"runtime"
	"os"
	"strings"
	"strconv"
)

func ForIt() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}
}

func ForPlay() {
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// Perfectly wrong code for a squareroot function
func Sqrt(x float64) float64 {
	a := 0
	if x < 0 {
		fmt.Println("Err: Error!")
	}

	for z := 1; z*z < int(x); z++ {
		return float64(z)
	}
	return float64(a)
}

func SwitcherTheWitcher() {
	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	}
}

func DayMotivator () {
	input := os.Args[1]
	dinput := strings.ToLower(input)
	
	switch dinput{
	case "monday":
		fmt.Println("Go hard this week")
	case "tuesday":
		fmt.Println("Don't slack, people who fail relent")
	case "wednesday":
		fmt.Println("You have done well, the results will show")

	}
}

func ageDillemma(){
	fmt.Println("Insert a number")

	if len(os.Args)<2{
		fmt.Println("Please provide a time argument (e.g. go run . 27)")
		return
	}

	input := os.Args[1]

	input2, err := strconv.Atoi(input)
	if err != nil{
		fmt.Println("Not a valid number")
		return
		}

	switch{
	case input2 < 12:
		fmt.Println("Good morning!")
	case input2 < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good evening")
	}
}
