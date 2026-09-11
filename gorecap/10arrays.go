package main

import (
	"fmt"
)

func Arraying() {
	// var array1 = [length]datatype{values}
	// var arrat2 = [...]datatype{values}
	array1 := [5]int{1, 2, 3}
	array2 := [...]int{4, 5, 6, 7, 8}
	array3 := [5]int{}
	array4 := [5]int{9, 10, 11, 12, 13}
	array5 := [5]int{1: 14, 2: 15}
	cars := [4]string{"Volvo", "BMW", "Ford", "Mazda"}

	var a [3]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a[1],a[2])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Println(array4)
	fmt.Println(array5)
	fmt.Println(cars)
	fmt.Println(cars[0])
	fmt.Println(array1[1])
	cars[0] = "Toyota"
	fmt.Println(cars)
	fmt.Println(len(cars))
	fmt.Println(len(array2))
}
