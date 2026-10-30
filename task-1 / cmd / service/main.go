package main

import "fmt"

func main() {
	var firstOperand, secondOperand int

	_, err := fmt.Scan(&firstOperand)
	if err != nil {
		return
	}

	_, err = fmt.Scan(&secondOperand)
	if err != nil {
		return
	}

	var operation string

	_, err = fmt.Scan(&operation)
	if err != nil {
		return
	}

	switch operation {
	case "+":
		fmt.Println(firstOperand + secondOperand)
	case "-":
		fmt.Println(firstOperand - secondOperand)
	case "*":
		fmt.Println(firstOperand * secondOperand)
	case "/":
		if secondOperand == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(firstOperand / secondOperand)
	default:
		fmt.Println("Invalid operation")
	}
}
