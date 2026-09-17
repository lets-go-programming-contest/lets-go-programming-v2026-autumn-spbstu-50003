package main

import "fmt"
func main() {
	var oper1 int
	var oper2 int
	var symb string
	_, err1 := fmt.Scanln(&oper1)
	if err1 != nil{
		fmt.Println("Invalid first operand")
		return
	}
	_, err2 := fmt.Scanln(&oper2)
	if err2 != nil{
		fmt.Println("Invalid second operand")
		return
	}
	fmt.Scanln(&symb)
	switch symb{
	case "+":
		fmt.Println(oper1 + oper2)
	case "-":
		fmt.Println(oper1 - oper2)
	case "*":
		fmt.Println(oper1 * oper2)
	case "/":
		if (oper2 != 0){
			fmt.Println(oper1 / oper2)
		}else{
			fmt.Println("Division by zero")
			return
		}
	default:
		fmt.Println("Invalid operation")
		return
	}
}