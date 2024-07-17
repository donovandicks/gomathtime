# gomathtime

Simple program to do simple math problems in the terminal
with a time limit.

Usage:
```
go run main.go [-s int] [-n int] [-o (multiplication, addition, subtraction, random)]

Usage of main:
  -n int
    	the maximum number the system can prompt for addition (default 10)
  -o string
    	the math operation to ask the user. defaults to a random operation for each problem (default "random")
  -s int
    	number of seconds to wait for an answer before quitting (default 3)
```
