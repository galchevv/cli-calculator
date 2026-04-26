package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Использование: ", os.Args[0], "<число> <операция> <число>")
		os.Exit(1)
	}

	a, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Printf("left operand parse: %v\n", err)
		os.Exit(1)
	}

	op := os.Args[2]

	b, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Printf("right operand parse: %v\n", err)
		os.Exit(1)
	}

	switch op {
	case "+":
		fmt.Printf("Сумма: %.2f", a+b)
	case "-":
		fmt.Printf("Разность: %.2f", a-b)
	case "*":
		fmt.Printf("Умножение: %.2f", a*b)
	case "/":
		fmt.Printf("Частное: %.2f", a/b)
	default:
		fmt.Println("Некорректная операция. Попробуйте снова...")
	}
}
