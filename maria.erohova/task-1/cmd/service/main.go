package main

import (
	"fmt"
)

func main() {
	var (
		o1, o2 int
		operation string
	)
	_, errO1 := fmt.Scanln(&o1)
	if errO1 != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, errO2 := fmt.Scanln(&o2)
	if errO2 != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, errOP := fmt.Scanln(&operation)
	if errOP != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch operation {
	case "+":
		fmt.Println(o1 + o2)
	case "-":
		fmt.Println(o1 - o2)
	case "*":
		fmt.Println(o1 * o2)
	case "/":
		if o2 != 0 {
			fmt.Println(o1 / o2)
		} else {
			fmt.Println("Division by zero")
			return
		}
	default:
		fmt.Println("Invalid operation")
		return
	}
}
