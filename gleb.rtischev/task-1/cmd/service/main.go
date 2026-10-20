package main

import (
	"fmt"
)

func main() {
	var (
		operation string
		a, b      int
	)

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	if _, err := fmt.Scan(&b); err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	if _, err := fmt.Scan(&operation); err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch operation {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b != 0 {
			fmt.Println(a / b)
		} else {
			fmt.Println("Division by zero")
		}
	default:
		fmt.Println("Invalid operation")
	}
}
