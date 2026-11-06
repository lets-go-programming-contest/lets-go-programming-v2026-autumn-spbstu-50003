package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		departments int
		staff       int
		degree      int
		operation   string
	)

	if _, err := fmt.Scan(&departments); err != nil {
		fmt.Fprintln(os.Stderr, "Invalid input for departments")

		return
	}

	for range departments {
		if _, err := fmt.Scan(&staff); err != nil {
			return
		}

		minTemp := 15
		maxTemp := 30

		for range staff {
			if _, err := fmt.Scan(&operation, &degree); err != nil {
				fmt.Fprintln(os.Stderr, "Invalid input for operation/degree")

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
