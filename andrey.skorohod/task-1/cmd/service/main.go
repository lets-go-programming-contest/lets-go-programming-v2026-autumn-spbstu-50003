package main

import (
	"fmt"
)

func main() {
	var (
		loperand, roperand int
		operation          string
	)

	_, errLop := fmt.Scanln(&loperand)
	if errLop != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, errRop := fmt.Scanln(&roperand)
	if errRop != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, errOp := fmt.Scanln(&operation)
	if errOp != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch operation {
	case "+":
		fmt.Println(loperand + roperand)

	case "-":
		fmt.Println(loperand - roperand)

	case "*":
		fmt.Println(loperand * roperand)

	case "/":
		if roperand == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(loperand / roperand)

	default:
		fmt.Println("Invalid operation")
		return
	}
}
