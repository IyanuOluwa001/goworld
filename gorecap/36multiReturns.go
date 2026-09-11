package main

func multipleRet(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}
