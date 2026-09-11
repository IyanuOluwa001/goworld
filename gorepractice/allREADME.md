# Lessons

When writing a test, we are testing a functions behaviour when given inputs and outputs.

We want to test:
- input
- edge cases
- banner correctness

What to check:
- Did my program run correctly?
- Does my function return the correct string?

Struture of a Testfile is:
Naming: asciiArt_test.go
Function Declaration:

```go
func TestAscii (t *testing.T)
```

```go
package main
//because the function I am testing is also in the main package. If the function I am testing happens to be in another package, then I have to use that package.
import "testing" 

func TestAsciiSingleLetter (t *testing.T){
    result, err:= asciiArt("A", "standard")

    if err !=nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if result == ""{
        t.Fatalf("expected output, empty string")
    }
}
```

To run a test, we use:
```go
go test -v
```

###Recall:

```go
package main

import "testing"

func TestAsciiArt (t *testing.T){
    result, err:= asciiArt("A", "standard")

    if err != nil{
        t.Fatalf("unexpected error: %v", err)
    }

    if result==""{
        t.Fatalf("expected output, got empty string")
    }
}
```
## Recode1:
```go
package main

import "testing"

func TestAsciiArt(t *testing.T){
    result,err := asciiArt("A", "standard")

    if err!=nil{
        t.Fatalf("unexpected error: %v, err")
    }
    expected := `EXACT ASCII OUTPUT HERE`

    if result!=expected{
        t.Errorf("got:\n%v\n\nwant:\n%v", result, expected)
        //t.Errorf("\nGOT:\n%s\n\nWANT:\n%s\n", result, expected)
    }

}
```

## Recode2:
```go
package main

import "testing"

func TestAsciiArt(t *testing.T){

    tests:= []struct{
        name string
        input string
        banner string
        expected string
    }{
        {
            name: "single valid input",
            input: "A",
            banner: "standard"
            expected: `ASCII OUTPUT`,
        },
        {
            name: "empty input",
            input: "",
            banner: "standard",
            expected: "",
        },
        {
            name: "new line",
            input: "\n",
            banner: "standard",
            expected: `ASCII OUTPUT`,
        },
    }

    for _, tt:= range tests{

        results, err:= asciiArt(tt.input, tt.banner)

        if err!= nil{
            t.Fatalf("unexpected error: %v", err)
        }

        if result != tt.expected{
            t.Errorf(
                "\nTEST: %s\nGOT:\n%s\n\nWANT:\n%s\n",
                tt.name,
                result,
                tt.expected,
            )
        }
    }
}
```
### Note:
- bool, int, uint32, float64, string, byte, rune, complex128
- In go, we value simplicity, readability, and scalability over fancy abstractions.

