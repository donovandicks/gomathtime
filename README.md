# gomathtime

Simple program to do simple math problems in the terminal
with a time limit.

Usage:
```
go run main.go [-s int] [-n int] [-o (multiplication, addition, subtraction, random)]

Usage of main:
  -n int
    	the maximum number the system can prompt for calculations (default 10)
  -o string
    	the math operation to ask the user. defaults to a random operation for each problem (default "random")
  -s int
    	number of seconds to wait for an answer before quitting (default 3)
```

Example:
```
go run main.go -s 3 -n 10 -o addition

Teacher: add 6 and 3
User: 9
Teacher: Correct!
Teacher: add 1 and 3
User: 4
Teacher: Correct!
Teacher: add 1 and 2
User: 3
Teacher: Correct!
Teacher: add 5 and 6
User:
Teacher: Oops! 3 seconds have passed. Game over!
```
