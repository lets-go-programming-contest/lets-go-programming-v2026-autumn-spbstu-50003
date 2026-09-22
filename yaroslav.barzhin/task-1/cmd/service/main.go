package main

import "fmt"

func main() {
	var (
		left, right, result float64
		op                  string
	)

	if _, err := fmt.Scan(&left); err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	if _, err := fmt.Scan(&right); err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	if _, err := fmt.Scan(&op); err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch op {
	case "+":
		result = left + right
	case "-":
		result = left - right
	case "*":
		result = left * right
	case "/":
		if right == 0 {
			fmt.Println("Division by zero")
			return
		}
		result = left / right
	default:
		fmt.Println("Invalid operation")
		return
	}
	fmt.Println(result)
}
