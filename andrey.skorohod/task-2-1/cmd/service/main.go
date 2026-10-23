package main

import (
	"fmt"
)

const (
	highBord = 30
	lowBord  = 15
)

type Range struct {
	Min int
	Max int
}

func (tempRange *Range) Update(sign string, temp int) int {
	if sign == ">=" && temp > tempRange.Min {
		tempRange.Min = temp
	} else if sign == "<=" && temp < tempRange.Max {
		tempRange.Max = temp
	}

	if tempRange.Min > tempRange.Max {
		return -1
	}

	return tempRange.Min
}

func main() {
	var departments, employees int
	if _, err := fmt.Scan(&departments); err != nil || departments < 1 || departments > 1000 {
		fmt.Println("Invalid number of departments: ", err)

		return
	}

	for range departments {
		if _, err := fmt.Scan(&employees); err != nil || employees < 1 || employees > 1000 {
			fmt.Println("Invalid number of employees: ", err)

			return
		}

		tempRng := Range{Min: lowBord, Max: highBord}

		for range employees {
			var (
				sign string
				temp int
			)

			if _, err := fmt.Scan(&sign, &temp); err != nil || (sign != ">=" && sign != "<=") {
				fmt.Println("Invalid sign or temperature value: ", err)

				return
			}

			fmt.Println(tempRng.Update(sign, temp))
		}
	}
}
