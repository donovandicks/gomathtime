package main

import (
	"math/rand/v2"

	"golang.org/x/exp/maps"
)

var (
	operations map[string]Operator = map[string]Operator{
		"addition":       {verb: "add", fn: Add},
		"substraction":   {verb: "subtract", fn: Sub},
		"multiplication": {verb: "multiply", fn: Mul},
	}
)

type Operator struct {
	verb string
	fn   func(num1, num2 int) int
}

func Add(num1, num2 int) int { return num1 + num2 }
func Sub(num1, num2 int) int { return num1 - num2 }
func Mul(num1, num2 int) int { return num1 * num2 }

// getOperator returns the math function associated with the given
// operator and the English verb for the operation.
//
// If the given operator is "random", it will select one of the
// operations at random.
func getOperator(name string) *Operator {
	if name == "random" {
		operatorNames := maps.Keys(operations)
		name = operatorNames[rand.IntN(len(operatorNames))]
	}

	op := operations[name]
	return &op
}
