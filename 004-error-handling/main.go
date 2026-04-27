package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintln(os.Stderr, "usage: ", os.Args[0], "<left operand> <operation> <right operand>")
		os.Exit(1)
	}

	a, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "left operand parse:", err)
		os.Exit(1)
	}

	op := os.Args[2]

	b, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "right operand parse:", err)
		os.Exit(1)
	}

	result, err := calc(op, a, b)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Printf("result: %.2f", result)
}

func calc(op string, a, b float64) (float64, error) {
	var result float64
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		result = a / b
	default:
		return 0, errors.New("incorrect operation, try again")
	}

	return result, nil
}
