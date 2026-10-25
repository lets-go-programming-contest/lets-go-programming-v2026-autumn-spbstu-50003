package main

import (
	"fmt"
)

func main() {
	var (
		departments int
		staff       int
		degree      int
		operation   string
	)
	if _, err := fmt.Scan(&departments); err != nil {
		fmt.Println("Invalid input for departments")
		return
	}
	for i := 0; i < departments; i++ {
		if _, err := fmt.Scan(&staff); err != nil {
			return
		}
		minTemp := 15
		maxTemp := 30
		for j := 0; j < staff; j++ {
			if _, err := fmt.Scan(&operation, &degree); err != nil {
				fmt.Println("Invalid input for operation/degree")
				return
			}
			if operation == ">=" {
				if degree > minTemp {
					minTemp = degree
				}
			} else if operation == "<=" {
				if degree < maxTemp {
					maxTemp = degree
				}
			}
			if minTemp <= maxTemp {
				fmt.Println(minTemp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
