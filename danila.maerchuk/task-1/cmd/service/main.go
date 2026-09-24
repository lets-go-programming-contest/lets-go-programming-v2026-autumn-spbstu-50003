package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	text1, _ := reader.ReadString('\n')
	text1 = strings.TrimSpace(text1)
	a, err1 := strconv.Atoi(text1)
	if err1 != nil {
		fmt.Println("Invalid first operand")
		return
	}

	text2, _ := reader.ReadString('\n')
	text2 = strings.TrimSpace(text2)
	b, err2 := strconv.Atoi(text2)
	if err2 != nil {
		fmt.Println("Invalid second operand")
		return
	}

	op, _ := reader.ReadString('\n')
	op = strings.TrimSpace(op)

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
