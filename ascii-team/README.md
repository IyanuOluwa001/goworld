# ASCII Art Generator (Standard Banner)

## Overview

This project is a **Go command-line application** that converts normal
text into **ASCII art** using the `standard.txt` banner format.

ASCII art represents characters using patterns of symbols arranged
across multiple rows to visually form letters.

Example:

Input:

    Hello

Output:

     _   _      _ _
    | | | | ___| | | ___
    | |_| |/ _ \ | |/ _ \
    |  _  |  __/ | | (_) |
    |_| |_|\___|_|_|\___/

------------------------------------------------------------------------

# Project Purpose

This project demonstrates:

-   Reading files in Go
-   Handling command-line arguments
-   Manipulating strings
-   Mapping ASCII characters to visual patterns
-   Writing basic automated tests

------------------------------------------------------------------------

# How the Program Works

The program follows these steps:

1.  Accepts a **text input argument** from the command line.
2.  Reads the ASCII banner file (`standard.txt`).
3.  Builds a **character map** for printable ASCII characters.
4.  Converts the input text into ASCII art using the banner patterns.
5.  Prints the ASCII art to the terminal.

------------------------------------------------------------------------

# Supported ASCII Characters

The program supports all **printable ASCII characters**:

    ASCII 32 → ASCII 126

This range contains:

    95 printable characters

Examples include:

-   Space
-   Letters (A--Z, a--z)
-   Numbers (0--9)
-   Symbols (!@#\$%\^&\* etc.)

------------------------------------------------------------------------

# Command-Line Arguments

The program expects **two command-line arguments**:

    go run . "text"

### Argument Breakdown

  Argument   Description
  ---------- --------------------------------------------------
  `"text"`   The string that will be converted into ASCII art

Example:

    go run . "Hello"

------------------------------------------------------------------------

# Newline Support

The program supports **multi-line text** using `\n`.

Example:

    go run . "Hello\nWorld"

Output:

    (ASCII art for Hello)
    (ASCII art for World)

------------------------------------------------------------------------

# Project Structure

Example folder layout:

    ascii-art/
    │
    ├── main.go
    ├── main_test.go
    └── standard.txt

### File Descriptions

  File             Purpose
  ---------------- -----------------------------------------
  `main.go`        Main program containing ASCII art logic
  `main_test.go`   Tests validating banner file structure
  `standard.txt`   ASCII art banner definitions

------------------------------------------------------------------------

# Banner File Format

The banner file defines how each ASCII character is drawn.

Each character contains:

    8 rows of ASCII art
    + 1 empty separator line

Total characters:

    95 characters

Total expected lines:

    95 × 9 = 855 lines

------------------------------------------------------------------------

# Running the Program

### Step 1 --- Navigate to the project directory

    cd ascii-art

### Step 2 --- Run the program

    go run . "Hello"

Example:

    go run . "ASCII"

------------------------------------------------------------------------

# Running Tests

The project includes tests to validate the structure of the banner file.

Run:

    go test

Example output:

    PASS
    ok   ascii-art   0.001s

------------------------------------------------------------------------

# Tests Included

The test file validates:

  Test                    Purpose
  ----------------------- ----------------------------------------
  Banner file exists      Ensures `standard.txt` can be opened
  Banner file not empty   Prevents blank banner errors
  Banner line count       Ensures ASCII banner format is correct
  ASCII printable range   Validates ASCII character range

------------------------------------------------------------------------

# Error Handling

The program handles several common errors.

### Incorrect number of arguments

    Usage: go run . "text"

### Missing banner file

    Failed to read standard.txt

### Invalid banner format

Detected when running:

    go test

------------------------------------------------------------------------

# Example Usage

### Example 1

Command:

    go run . "Hi"

Output:

    (ASCII representation of Hi)

------------------------------------------------------------------------

### Example 2 (Multi-line)

Command:

    go run . "Hello\nWorld"

Output:

    ASCII art for Hello
    ASCII art for World

------------------------------------------------------------------------

# Limitations

-   Only the **standard banner format** is supported.
-   Only **printable ASCII characters (32--126)** are supported.
-   The file `standard.txt` must exist in the project directory.

------------------------------------------------------------------------

# Future Improvements

Possible improvements include:

-   Supporting multiple banner styles (`shadow`, `thinkertoy`, etc.)
-   Writing output to a file
-   Better error messages for unsupported characters
-   Performance improvements for large input text

------------------------------------------------------------------------

# Technologies Used

-   **Go programming language**
-   **Go testing package**

Standard libraries used:

    os
    strings
    testing

------------------------------------------------------------------------

# Learning Outcomes

After completing this project, a developer should understand:

-   How Go CLI applications work
-   How to process command-line input
-   How to read files in Go
-   How to manipulate strings
-   How to write automated tests

------------------------------------------------------------------------

# Author

L2E ASCII Group work code using **Go**.
