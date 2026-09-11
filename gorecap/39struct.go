package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func strucInfo() {
	var employee1 Person
	var employee2 Person

	// employee1 information
	employee1.name = "Hege"
	employee1.age = 45
	employee1.job = "Technical Director"
	employee1.salary = 3500

	// employee2 information
	employee2.name = "Ajoke"
	employee2.age = 32
	employee2.job = "Frontend Developer"
	employee2.salary = 2300

	// Access and print employee1 info
	fmt.Println("Name:", employee1.name)
	fmt.Println("Age:", employee1.age)
	fmt.Println("Job:", employee1.job)
	fmt.Println("Salary", employee1.salary)
	fmt.Println()

	// Access and print employee2 info
	fmt.Println("Name:", employee2.name)
	fmt.Println("Age:", employee2.age)
	fmt.Println("Job:", employee2.job)
	fmt.Println("Salary:", employee2.salary)
}
