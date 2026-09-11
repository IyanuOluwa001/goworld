# ASCII Art Justify

A Go program that converts text into ASCII art using different banner styles and alignment options.
This project supports **left**, **right**, **center**, and **justify** text alignment directly in the terminal.

---

## Features

* Convert normal text into ASCII art
* Multiple banner styles:

  * `standard`
  * `shadow`
  * `thinkertoy`
* Alignment support:

  * `left` (default)
  * `right`
  * `center`
  * `justify`
* Automatically detects terminal width for alignment handling
* Handles multi-word justification spacing

---

## Project Structure

```bash
ascii-art-justify/
│
├── main.go
│
├── asciifunctions/
│   ├── asciiArt.go
│   └── justify.go
│
├── standard.txt
├── shadow.txt
├── thinkertoy.txt
│
└── README.md
```

---

# How It Works

The program:

1. Reads user input from the command line
2. Loads the selected ASCII banner file
3. Converts each character into ASCII art
4. Applies the selected alignment based on terminal width
5. Prints the formatted ASCII art output

---

# Installation

## Clone the Repository

```bash
git clone <repository-url>
cd ascii-art-justify
```

---

# Usage

```bash
go run . [OPTION] [STRING] [BANNER]
```

## Example

```bash
go run . --align=right "Hello World" standard
```

---

# Alignment Options

| Alignment | Description                             |
| --------- | --------------------------------------- |
| left      | Default alignment                       |
| right     | Aligns text to the right side           |
| center    | Centers the ASCII art                   |
| justify   | Distributes spaces evenly between words |

---

# Banner Options

| Banner     | Description             |
| ---------- | ----------------------- |
| standard   | Standard ASCII font     |
| shadow     | Shadow-styled ASCII art |
| thinkertoy | Minimal decorative font |

---

# Examples

## Left Alignment (Default)

```bash
go run . "Hello" standard
```

---

## Right Alignment

```bash
go run . --align=right "Hello" shadow
```

---

## Center Alignment

```bash
go run . --align=center "Hello World" thinkertoy
```

---

## Justify Alignment

```bash
go run . --align=justify "Go is awesome" standard
```

---

# Error Handling

The program validates:

* Invalid flags
* Invalid alignment types
* Invalid banner names
* Incorrect number of arguments
* Banner file reading errors

Example:

```bash
wrong aligntype
```

---

# Core Functions

## `AsciiArt()`

Located in:

```bash
asciifunctions/asciiArt.go
```

Responsibilities:

* Reads banner files
* Converts characters into ASCII art
* Builds output line by line
* Applies alignment

---

## `ApplyAlignment()`

Located in:

```bash
asciifunctions/justify.go
```

Responsibilities:

* Applies left, right, center, or justify alignment
* Calculates spacing using terminal width
* Handles word spacing for justification

---

## `GetTerminalWidth()`

Located in:

```bash
asciifunctions/justify.go
```

Responsibilities:

* Detects terminal width using:

```bash
stty size
```

* Used for dynamic alignment calculations

---

# Requirements

* Go 1.20+
* Unix/Linux/macOS terminal environment
* `stty` command available

---

# Notes

* The program assumes all banner files are formatted correctly.
* Justify alignment works best with multiple words.
* Terminal width affects alignment output.

---

# Future Improvements

* Windows terminal support
* Multi-line input handling
* Colored ASCII art
* Custom banner support
* Better error messages

---

# Author
Built with Go for learning terminal rendering, string manipulation, and text alignment algorithms by lodesanya, iadejare, and oolonibu
