1. install go through vscode extension
2. use code . to open a new folder directly from terminal
3. do shift ctrl P and
   paste:
   Go: Install/Update Tools
   and select all, then press ok
4. create go file, write your code and run:
    go mod tidy
    the run it with:
    go run .
5. Use // for single line comment and /* */ for multiline comment
6. variable in go include int, float32, string, bool
7. We use var to declare variable e.g:
    var variablename type = value
    var username string = "Moyo"
    var username2 = "Jane" //type is inferred
    x := 2 //type is inferred
8. We can also use 
    username := Moyo //here, go decides the type of variable based on the value provided
9. Go is strongly typed, you cannot add string with integer, integers must be converted
10. You cannot use := without assigning a value to it, and it can only be used inside a function, plus
    assignment of value cannot be done seperately

**int**
- is 32bit in 32bit systems and 64bit in 64 bit systems
- -2147483648 to 2137483647 in 32bit system
- -9223372036854775808 to 9223372036854775807 ub 64bit systems

**int8**
- 8bits/1 byte
- -128 to 127

**int16**
- 16 bits/ 2 byte
- -32768 to 32767

**int32**
- 32 bits/ 4 byte
- -2147483648 to 2147483647

**int64**
- 64 bits/ 8 byte
- -9223372036854775808 to 9223372036854775807

**Go Types**
To use some math function, you call: "math/complx"
- bool
- string
- int int8 int16 int32(rune) int63
- uint uint8 uint16 uint32 uint64 uintptr
- float32 float64
- complex64 complex128


