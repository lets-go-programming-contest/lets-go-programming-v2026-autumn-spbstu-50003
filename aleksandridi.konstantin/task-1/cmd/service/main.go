package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	firstOperand, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	scanner.Scan()
	secondOperand, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	scanner.Scan()
	op := scanner.Text()
	switch op {
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
