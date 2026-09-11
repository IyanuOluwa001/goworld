# ASCII Art Generator

## Overview

This program is a command-line ASCII art generator written in Go.

It converts regular text into ASCII art using banner font files such as:

* `standard.txt`
* `shadow.txt`
* `thinkertoy.txt`

The program can:

* Print ASCII art directly to the terminal
* Save generated ASCII art into a file
* Support multi-line input using `\n`

---

## Features

* Convert plain text to ASCII art
* Choose different banner styles
* Save output to a file using `--output`
* Handle multi-line text input: Meaning that you can run 
1. Two Argument:
e.g. go run . Hello

2. Threee Argument:
e.g. go run . Hello Thinkertoy

3. Four Arguments:
e.g. go run . --output=output.txt Hello Standard
---

## Project Structure

```text
.
> testdata
├── empty.go
├── newline.go
└──single_A.go
── main.go
── AsciiArt.go
── asciiArt_test.go
── go.mod
── standard.txt
── shadow.txt
── thinkertoy.txt
└── README.md
```

---

## Requirements

* Go 1.18 or later

---

## Installation

Clone the repository:

```bash
git clone https://acad.learn2earn.ng/git/iadejare/ascii-art-output.git
cd ascii-art-output
```

---

## Usage

### Basic Syntax

```bash
go run . [OPTION] [STRING] [BANNER]
```

---

## Arguments

| Argument   | Description                                       |
| ---------- | ------------------------------------------------- |
| `[OPTION]` | Optional output flag                              |
| `[STRING]` | Text to convert into ASCII art                    |
| `[BANNER]` | Banner style (`standard`, `shadow`, `thinkertoy`) |

---

## Examples

### 1. Print using default banner (`standard`)

```bash
go run . "Hello"
```

---

### 2. Print using a custom banner

```bash
go run . "Hello" shadow
```

---

### 3. Save output to a file

```bash
go run . --output=result.txt "Hello" standard
```

This writes the ASCII art into:

```text
result.txt
```

---

### 4. Multi-line input

```bash
go run . "Hello\nWorld"
```

Output:

```text
(ASCII art for Hello)
(ASCII art for World)
```

---

## Supported Banner Files

The program expects banner files to exist in the project directory.

Examples:

* `standard.txt`
* `shadow.txt`
* `thinkertoy.txt`

The selected banner is loaded automatically:

```go
banner + ".txt"
```

---

## How It Works

The program:

1. Reads command-line arguments
2. Validates optional output flag
3. Loads the selected banner file
4. Converts each input character into ASCII art
5. Builds the output using `strings.Builder`
6. Prints or writes the result to a file

---

## Output Flag

To save output:

```bash
--output=<filename>
```

Example:

```bash
go run . --output=ascii.txt "Go" standard
```

---

## Error Handling

The program handles:

### Invalid number of arguments

```text
Usage: go run . [OPTION] [STRING] [BANNER]
```

---

### Invalid output flag

```text
Invalid flag usage
```

---

### Missing banner file

```text
go run . [STRING] [BANNER]
```

---

### File write failure

```text
Error Writing File
```

---

## Example Workflow

Generate and save ASCII art:

```bash
go run . --output=hello.txt "Hello World" shadow
```

Then view:

```bash
cat hello.txt
```

---

## Technologies Used

* Go
* Standard library packages:

  * `fmt`
  * `os`
  * `strings`

---

## Author

ASCII Art Generator built in Go as a command-line text rendering project.
