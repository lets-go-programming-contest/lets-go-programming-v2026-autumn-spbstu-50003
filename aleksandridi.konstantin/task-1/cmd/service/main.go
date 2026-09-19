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
	num1, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	scanner.Scan()
	num2, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	scanner.Scan()
	op := scanner.Text()
	switch op {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(num1 / num2)
	default:
		fmt.Println("Invalid operation")
	}
}
