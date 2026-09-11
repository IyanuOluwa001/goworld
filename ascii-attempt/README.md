# ASCII Art Generator (Go)

## Overview

This project is a **Go program that converts normal text into large ASCII-art text** using predefined banner templates.

Instead of printing characters normally, the program prints a **graphic representation made of ASCII symbols**.

Example:

Input

```
Hello
```

Output (example style)

```
 _    _          _   _
| |  | |        | | | |
| |__| |   ___  | | | |   ___
|  __  |  / _ \ | | | |  / _ \
| |  | | |  __/ | | | | | (_) |
|_|  |_|  \___| |_| |_|  \___|
```

The ASCII design is generated from **banner template files**.
Basically my
├── standard.txt
├── shadow.txt
└── thinkertoy.txt
files.
---

# My Project Structure

```
ascii-art/
│
├── main.go
├── main_test.go
│
├── standard.txt
├── shadow.txt
└── thinkertoy.txt
│
└── README.md
```

### File Description

| File             | Purpose                         |
| ---------------- | ------------------------------- |
| `main.go`        | Contains the main program logic |
| `main_test.go`   | Unit tests for key functions    |
| `standard.txt`   | ASCII banner template           |
| `shadow.txt`     | ASCII banner template           |
| `thinkertoy.txt` | ASCII banner template           |
| `README.md`      | Project documentation           |

---

# Requirements

Before running the project you need:

* Go installed (version 1.20 or newer recommended)

Check installation:

```
go version
```

Example output:

```
go version go1.22 linux/amd64
```

If Go is not installed, download it from:

https://go.dev/dl/

---

# How the Program Works

The program follows these steps:

1. Accepts **text input from the command line**
2. Loads a **banner template file**
3. Converts each character into **ASCII-art**
4. Prints the final output to the terminal

ASCII characters are stored inside the banner files.

Each character is represented by **8 lines of ASCII art**.

Example concept:

```
H = 8 lines
e = 8 lines
l = 8 lines
```

The program combines these rows together to form the final result.

---

# Running the Program

Basic usage:

```
go run . "text" template banner
```

### Example 1

```
go run . "Hello" standard
```

### Example 2

```
go run . "Hello" shadow
```

### Example 3

```
go run . "Hello" thinkertoy
```

---

# Using New Lines

The program supports newline characters.

Example:

```
go run . "Hello\nWorld" standard
```

Output:

```
$
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
$
$
__          __                 _       _  $
\ \        / /                | |     | | $
 \ \  /\  / /    ___    _ __  | |   __| | $
  \ \/  \/ /    / _ \  | '__| | |  / _` | $
   \  /\  /    | (_) | | |    | | | (_| | $
    \/  \/      \___/  |_|    |_|  \__,_| $
                                          $
```

---

# Supported Characters

The program supports **all printable ASCII characters**:

```
space
letters
numbers
punctuation
special symbols
```

ASCII range supported:

```
32 → 126
```

Examples:

```
A-Z
a-z
0-9
!@#$%^&*()
```

---

# Running Tests

The project includes unit tests.

Run them using:

```
go test
```

Example output:

```
PASS
ok      ascii-art       0.002s
```

Tests check:

* banner file reading
* ASCII map creation
* character height
* space character existence

---

# Example Output

Command:

```
go run . "Go" standard
```

Example output:

```
   ____  
  / ___| 
 | |  _  
 | |_| | 
  \____| 
```

---

# Common Errors

### Error: Wrong number of arguments

If you see:

```
Usage: go run . "text" banner
```

It means the command was not formatted correctly.

Correct format:

```
go run . "Hello" standard
```

---

### Error: Banner file not found

If you see an error like:

```
Error: open standard.txt: no such file or directory
```

Make sure the banner files exist in the project directory.

---

# Development Notes

The project is built using the **Go standard library only**.

Main packages used:

```
fmt
os
strings
testing
```

Core program functions:

| Function           | Purpose                    |
| ------------------ | -------------------------- |
| `readBannerFile()` | Reads banner template      |
| `buildAsciiMap()`  | Builds ASCII character map |
| `renderAscii()`    | Renders ASCII art output   |

---

# Learning Goals

This project helps developers learn:

* Go file handling
* Go maps
* slices
* string manipulation
* CLI argument parsing
* modular code design
* unit testing

---

# Example Developer Workflow

Clone or download the project:

```
git clone <repository>
cd ascii-art
```

Run the program:

```
go run . "Hello Developer" standard
```

Run tests:

```
go test
```

---

# License

This project is provided for educational purposes.
