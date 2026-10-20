package main

import "fmt"

func main() {
	var (
		firstInt  int
		secondInt int
		operand   string
	)
	_, err := fmt.Scan(&firstInt)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	_, err = fmt.Scan(&secondInt)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	_, err = fmt.Scan(&operand)
	if err != nil {
		fmt.Println("Invalid operand")
		return
	}
	switch operand {
	case "+":
		fmt.Println(firstInt + secondInt)
	case "-":
		fmt.Println(firstInt - secondInt)
	case "*":
		fmt.Println(firstInt * secondInt)
	case "/":
		if secondInt == 0 {
			fmt.Println("Division by zero")
		} else {
			fmt.Println(firstInt / secondInt)
		}
	default:
		fmt.Println("Invalid operation")
	}
}
