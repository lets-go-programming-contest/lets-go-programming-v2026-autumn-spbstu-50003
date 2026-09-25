package main

import "fmt"

func main() {
	var (
		a    int
		b    int
		symb string
	)
	_, errA := fmt.Scanln(&a)
	if errA != nil {
		fmt.Println("Invalid first operand")
		return
	}
	_, errB := fmt.Scanln(&b)
	if errB != nil {
		fmt.Println("Invalid second operand")
		return
	}
	_, errSymb := fmt.Scanln(&symb)
	if errSymb != nil {
		fmt.Println("Invalid operation")
		return
	}
	switch symb {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b != 0 {
			fmt.Println(a / b)
		} else {
			fmt.Println("Division by zero")
			return
		}
	default:
		fmt.Println("Invalid operation")
		return
	}
}
