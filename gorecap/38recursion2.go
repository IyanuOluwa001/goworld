package main

func factorial_recursion(x float64) (y float64) {
	// n! = n x (n-1) x (n-2) x ... x 1
	if x > 0 {
		y = x * factorial_recursion(x-1)
	} else {
		y = 1
	}
	// y = 4 x (4-1) x (4-2) x (4-3) = 24
	return // also known as naked return
}
