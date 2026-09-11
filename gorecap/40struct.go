package main

import (
	"fmt"
)

type Data struct {
	Name   string
	age    int
	job    string
	salary int
}

func structData() {
	var user1 Data
	var user2 Data

	// User1 Data
	user1.Name = "Stephen"
	user1.age = 27
	user1.job = "Academic writer"
	user1.salary = 2000000

	// User2 Data
	user2.Name = "James"
	user2.age = 23
	user2.job = "Freelancer"
	user2.salary = 200000

	// Print user1 data by calling a function
	printData(user1)

	// Print user2 data by calling a function
	printData(user2)
}

func printData(user Data) {
	fmt.Println("Name:", user.Name)
	fmt.Println("Age:", user.age)
	fmt.Println("Job:", user.job)
	fmt.Println("Salary:", user.salary)
	fmt.Println()
}

//Struct fields are accessed using a dot.
func struct2(){

type Vertex struct {
	X int
	Y int
}

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	
	v1 := Vertex{3, 6}
	p := &v1
	p.X = 1e9
	fmt.Println(v1)

	var (
	v2 = Vertex{1, 2}  // has type Vertex
	v3 = Vertex{X: 1}  // Y:0 is implicit
	v4 = Vertex{}      // X:0 and Y:0
	p2  = &Vertex{1, 2} // has type *Vertex
)
	fmt.Println(v2, p2, v3, v4)

}

