package main

import (
	"fmt"
)

func processEmployees(employeesCount int) {
	minTemp := 15
	maxTemp := 30

	for range employeesCount {
		var (
			operation string
			val       int
		)

		if _, err := fmt.Scan(&operation, &val); err != nil {
			break
		}

		if operation == ">=" {
			if val > minTemp {
				minTemp = val
			}
		} else if operation == "<=" {
			if val < maxTemp {
				maxTemp = val
			}
		}

		if minTemp <= maxTemp {
			fmt.Println(minTemp)
		} else {
			fmt.Println(-1)
		}
	}
}

func main() {
	var departmentNumber int
	if _, err := fmt.Scan(&departmentNumber); err != nil {
		return
	}

	if departmentNumber < 1 || departmentNumber > 1000 {
		return
	}

	for range departmentNumber {
		var employeesCount int
		if _, err := fmt.Scan(&employeesCount); err != nil {
			break
		}

		if employeesCount < 1 || employeesCount > 1000 {
			break
		}

		processEmployees(employeesCount)
	}
}
