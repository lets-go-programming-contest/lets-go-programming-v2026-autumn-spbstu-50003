package main

import (
	"fmt"
)

const (
	MinAllowedTemp = 15
	MaxAllowedTemp = 30
)

type TempRange struct {
	Min int
	Max int
}

func main() {
	var departments int
	if _, err := fmt.Scan(&departments); err != nil {
		fmt.Printf("Error reading departments: %v\n", err)
		return
	}

	for i := 0; i < departments; i++ {
		var employees int
		if _, err := fmt.Scan(&employees); err != nil {
			fmt.Println("Error reading employees")
			return
		}

		currentRange := TempRange{
			Min: MinAllowedTemp,
			Max: MaxAllowedTemp,
		}

		for j := 0; j < employees; j++ {
			var operator string
			var temp int
			if _, err := fmt.Scan(&operator, &temp); err != nil {
				fmt.Println("Error reading constraint")

				return
			}

			if operator == ">=" {
				if temp > currentRange.Min {
					currentRange.Min = temp
				}
			} else if operator == "<=" {
				if temp < currentRange.Max {
					currentRange.Max = temp
				}
			} else {
				fmt.Println("Invalid operator")
				return
			}

			if currentRange.Min > currentRange.Max {
				fmt.Println(-1)
			} else {
				fmt.Println(currentRange.Min)
			}
		}
	}
}
