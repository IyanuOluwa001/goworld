---

# ASCII Art Generator

A Go program that converts text into ASCII art using a custom font from `standard.txt`.

---

# Requirements

* `standard.txt` in the project folder

---

# Usage

Run the program with one argument:

```bash
go run main.go "Your Text Here"
```

Multiple lines:

```bash
go run main.go "Hello\nWorld"
```

---

# Project Structure

```
ascii-art-generator
│
├── main.go
├── main_test.go
├── standard.txt
```

---

# How It Works

1. Validate command line argument
2. Read `standard.txt`
3. Split input into lines and letters
4. Map each letter to its ASCII pattern
5. Print each line in ASCII art

---

# Notes

* Only ASCII characters 32–126 are supported
* Empty lines print automatically
* Output depends on `standard.txt`

---

# Technologies

* Go
* Standard libraries: `fmt`, `os`, `strings`

---

# Learning Outcomes

* File reading in Go
* String and slice processing
* ASCII mapping
* Command line handling
* Terminal output formatting

---
