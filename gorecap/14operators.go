package main

import (
	"fmt"
)

func OperAtor() {
	var (
		a = 4
		b = 5
	)
	c := a + b
	d := 10 + 15
	sum1 := 10 + 2
	sum2 := sum1 + 25
	sum3 := sum2 + sum2
	sum4 := sum2 - sum1
	sum5 := sum2 * a
	sum6 := sum2 / a
	sum7 := sum2 % 4
	sum2--
	a++        // b=a+1
	b += 5     // b=b+5
	c -= 6     // c=c-6
	d *= 2     // d=d*2
	sum1 /= 1  // sum1=sum1/1
	sum2 %= 5  // sum2=sum2%5
	sum3 &= 10 // sum3=sum3&10
	sum4 |= 11 // sum4=sum4|11
	sum5 ^= 12 // sum5=sum5^12

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(sum1)
	fmt.Println(sum2)
	fmt.Println(sum3)
	fmt.Println(sum4)
	fmt.Println(sum5)
	fmt.Println(sum6)
	fmt.Println(sum7)
}
