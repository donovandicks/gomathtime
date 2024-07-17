package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/maps"
)

const (
	DefaultMaxNumber   = 10
	DefaultWaitSeconds = 3
	DefaultOperator    = "random"
)

type Operator struct {
	verb string
	fn   func(num1, num2 int) int
}

var (
	waitSeconds  int
	maxNumber    int
	operatorName string

	operations map[string]Operator = map[string]Operator{
		"addition": {
			verb: "add",
			fn:   func(num1, num2 int) int { return num1 + num2 },
		},
		"substraction": {
			verb: "subtract",
			fn:   func(num1, num2 int) int { return num1 - num2 },
		},
		"multiplication": {
			verb: "multiply",
			fn:   func(num1, num2 int) int { return num1 * num2 },
		},
	}
)

// handleInput parses the user's provided answer and compares it
// to the expected answer.
//
// It returns a message to show the user and an indicator for
// whether the program should continue prompting the user.
func handleInput(input string, expected int) (string, bool) {
	answer, err := strconv.Atoi(strings.Trim(input, "\r\n"))
	if err != nil {
		return "Teacher: Please try again using a valid number", true
	}

	if answer == expected {
		return "Teacher: Correct!", true
	}

	return "Teacher: Wrong answer, game over!", false
}

// getOperator returns the math function associated with the given
// operator and the English verb for the operation.
//
// If the given operator is "random", it will select one of the
// operations at random.
func getOperator(name string) *Operator {
	var opName string
	if name == "random" {
		operatorNames := maps.Keys(operations)
		opName = operatorNames[rand.IntN(len(operatorNames))]
	} else {
		opName = name
	}

	op := operations[opName]
	return &op
}

// validateFlags checks that the given inputs are valid
func validateFlags() {
	if waitSeconds < 0 {
		log.Fatal("waitSeconds should be a positive integer\n")
	}

	operatorName = strings.ToLower(operatorName)
	_, valid := operations[operatorName]
	if !valid && operatorName != "random" {
		log.Fatalf(
			"Given operator %s is not valid, please choose one of ['addition', 'substraction', 'multiplication', 'random']", operatorName,
		)
	}
}

func main() {
	flag.IntVar(&waitSeconds, "s", DefaultWaitSeconds, "number of seconds to wait for an answer before quitting")
	flag.IntVar(&maxNumber, "n", DefaultMaxNumber, "the maximum number the system can prompt for calculations")
	flag.StringVar(&operatorName, "o", DefaultOperator, "the math operation to ask the user. defaults to a random operation for each problem")
	flag.Parse()

	validateFlags()
	waitTimeout := time.Duration(waitSeconds) * time.Second

	breakChan := make(chan struct{}, 1)
	inputChan := make(chan string)
	reader := bufio.NewReader(os.Stdin)

	for {
		num1, num2 := rand.IntN(maxNumber), rand.IntN(maxNumber)

		op := getOperator(operatorName)
		fmt.Printf("Teacher: %s %d and %d\nUser: ", op.verb, num1, num2)

		go func() {
			input, err := reader.ReadString('\n')
			if err != nil {
				if err.Error() == "EOF" {
					breakChan <- struct{}{}
				}
				log.Fatalf("failed to read user input: %v", err)
			}
			inputChan <- input
		}()

		select {
		case input := <-inputChan:
			msg, shouldContinue := handleInput(input, op.fn(num1, num2))
			fmt.Println(msg)
			if !shouldContinue {
				return
			}
		case <-breakChan:
			fmt.Println("\nTeacher: Goodbye!")
			return
		case <-time.After(waitTimeout):
			fmt.Printf("\nTeacher: Oops! %d seconds have passed. Game over!\n", waitSeconds)
			return
		}
	}
}
