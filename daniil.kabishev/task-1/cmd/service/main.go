package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	firstInput := strings.TrimSpace(scanner.Text())
	a, err := strconv.Atoi(firstInput)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	if !scanner.Scan() {
		return
	}
	secondInput := strings.TrimSpace(scanner.Text())
	b, err := strconv.Atoi(secondInput)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	if !scanner.Scan() {
		return
	}
	op := strings.TrimSpace(scanner.Text())

	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(a / b)
	default:
		fmt.Println("Invalid operation")
	}
}
