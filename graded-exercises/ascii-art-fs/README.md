# ASCII-ART-FS

### OVERVIEW
The ascii-art-fs mimics the function of ascii-arts by displaying different forms of banner texts based on what was passed into it. The output is generated based on the banner file selected.

### FEATURES
- It converts tests into ASCII art 
- Handles different arguments
- Provides optional support for additional flags
- Supports multiple banner templates


### INSTALLATION/SETUP
Clone the repository:
git clone https://acad.learn2earn.ng/git/ogajayi/ascii-art-fs

- Ensure Go is installed:
go version

---
### USAGE
The program must be run using:

go run . [STRING] [BANNER]
Example: go run. "hello" standard

Output: 

 _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  


### HANDLING ERROR 

- If incorrect arguments are provided, the program prints:

Usage: go run . [STRING] [BANNER]

### PROJECT STRUCTURE

├── main.go
│── standard.txt
│── shadow.txt
│── thinkertoy.txt
├── main_test.go/
└── README.md
