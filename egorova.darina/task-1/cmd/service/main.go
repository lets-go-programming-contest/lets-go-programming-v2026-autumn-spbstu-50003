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

	first, ok := readLine(scanner)
	if !ok {
		fmt.Println("Invalid first operand")
		return
	}

	a, err := strconv.Atoi(first)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	second, ok := readLine(scanner)
	if !ok {
		fmt.Println("Invalid second operand")
		return
	}

	b, err := strconv.Atoi(second)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	operation, ok := readLine(scanner)
	if !ok || len(operation) != 1 {
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
		if b == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(a / b)
	default:
		fmt.Println("Invalid operation")
	}
}

func readLine(scanner *bufio.Scanner) (string, bool) {
	if !scanner.Scan() {
		return "", false
	}
	return strings.TrimSpace(scanner.Text()), true
}
