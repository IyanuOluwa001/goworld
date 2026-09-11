package main

import (
	"fmt"
)

func ForVer() {
	i := 15
	var l int = 18
	j := "digger"
	var k float32 = 3.141
	var txt string = "Hello World! again."
	a := true
	var b bool = false
	var c int = 4132
	var d int = -1243
	var e float32 = 123.78
	var f float32 = 3.4e+38
	var g float64 = 1.7e+300
	var balance uint = 500

	fmt.Printf("%v\n", i)   // Prints the value of the argument
	fmt.Printf("%#v\n", i)  // Prints the value in Go-syntax format
	fmt.Printf("%T\n", i)   // Prints the data type
	fmt.Printf("%v%%\n", l) // Prints the value and a % sign
	fmt.Println()

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)
	fmt.Println()

	fmt.Printf("%b\n", i)   // base2
	fmt.Printf("%d\n", i)   // base10
	fmt.Printf("%+d\n", i)  // base10 and always show sign
	fmt.Printf("%o\n", i)   // base8
	fmt.Printf("%O\n", i)   // base8, with leading Oo
	fmt.Printf("%x\n", i)   // base16, lowercase
	fmt.Printf("%X\n", i)   // base16, uppercase
	fmt.Printf("%#x\n", i)  // base16, with leading 0x
	fmt.Printf("%4d\n", i)  // Pad with spaces (width 4, right justified)
	fmt.Printf("%-4d\n", i) // Pad with spaces (width 4, left justified)
	fmt.Printf("%04d\n", i) // Pad with zeroes (width 4)
	fmt.Println()

	fmt.Printf("%s\n", j)   // Prints the value as plain string
	fmt.Printf("%q\n", i)   // Prints the value as a double-quoted string
	fmt.Printf("%8s\n", j)  // Prints the value as plain string (width 8, right justified)
	fmt.Printf("%-8s\n", j) // Prints the value as plain string (width 8, left justified)
	fmt.Printf("%X\n", i)   // Prints the value as hex dump of byte values
	fmt.Printf("%x\n", i)   // Prints the value as hex dump with spaces
	fmt.Println()

	fmt.Printf("%t\n", a)
	fmt.Printf("%t\n", b)
	fmt.Println()

	fmt.Printf("%e\n", k)
	fmt.Printf("%f\n", k)
	fmt.Printf("%.2f\n", k)
	fmt.Printf("%6.2f\n", k)
	fmt.Printf("%g\n", k)
	fmt.Println()

	/*
		Signed integers: e.g int, can store both neg and pos values
		Unsigned integers: e.g uint, can only store non-negative values
	*/

	fmt.Printf("Type: %T, Value:%v\n", c, c)
	fmt.Printf("Type: %T, Value: %v\n", d, d)
	fmt.Println()

	fmt.Println(balance)
	fmt.Printf("Type: %T, value: %v\n", e, e)
	fmt.Printf("Type: %T, value: %v\n", f, f)
	fmt.Printf("Type: %T, value: %v\n", g, g)

	/*
			INTEGERS
		int: can store positive and negative and is called signed
		- is 32bit in 32bit systems and 64bit in 64 bit systems
		- -2147483648 to 2137483647 in 32bit system
		- -9223372036854775808 to 9223372036854775807 ub 64bit systems

		int8
		- 8bits/1 byte
		- -128 to 127

		int16
		- 16 bits/ 2 byte
		- -32768 to 32767

		int32
		- 32 bits/ 4 byte
		- -2147483648 to 2147483647

		int64
		- 64 bits/ 8 byte
		- -9223372036854775808 to 9223372036854775807

		uint: can only be positive and called unsigned
		- is 32bit in 32bit systems and 64bit in 64 bit systems
		- 0 to 4294967295 in 32bit system
		- 0 to 18446744073709551615 is 64bit systems

		int8
		- 8bits/1 byte
		- 0 to 255

		int16
		- 16 bits/ 2 byte
		- 0 to 65535

		int32
		- 32 bits/ 4 byte
		- 0 to 4294967295

		int64
		- 64 bits/ 8 byte
		- 0 to 18446744073709551615

		FLOATS
		Floats accepts data types with decimal points, whether positive or negative
		float32 (32 bits): -3.4e+38 to 3.4e+38
		float64 (64 bits): -1.7e+308 to 1.7e+308 - this is default if 32 or 64 is not determined for float
	*/
}
