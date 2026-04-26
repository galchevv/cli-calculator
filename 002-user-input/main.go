package main

import (
	"fmt"
)

func main() {
	var a, b float64
	var op string
	fmt.Println("Введите числа")
	fmt.Scan(&a, &b)
	fmt.Println(a, b)
	fmt.Println("Введите операцию: +, -, *, /")
	fmt.Scan(&op)

	switch op {
	case "+":
		fmt.Println("Сумма:", a+b)
	case "-":
		fmt.Println("Разность:", a-b)
	case "*":
		fmt.Println("Произведение:", a*b)
	case "/":
		fmt.Println("Частное:", a/b)
	default:
		fmt.Println("Операция не определена!")
	}
}
