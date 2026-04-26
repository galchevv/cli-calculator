package main

import "fmt"

func main() {
	x1 := 20
	x2 := 6
	s := x1 + x2
	d := x1 - x2
	m := x1 * x2
	f := x1 / x2
	r := x1 % x2
	fmt.Println("Первое число", x1)
	fmt.Println("Второе число", x2)
	fmt.Println("Сумма", s)
	fmt.Println("Разность", d)
	fmt.Println("Произведение", m)
	fmt.Println("Частное", f)
	fmt.Println("Остаток", r)
}
