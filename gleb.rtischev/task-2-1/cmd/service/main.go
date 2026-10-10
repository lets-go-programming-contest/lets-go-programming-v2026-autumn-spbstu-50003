package main

import (
	"fmt"
)

type temperature struct {
	min int
	max int
}

func main() {
	var (
		departments, staff, degree int
		operation                  string
	)

	if _, err := fmt.Scan(&departments); err != nil {
		return
	}

	for range departments {
		temp := temperature{15, 30}

		if _, err := fmt.Scan(&staff); err != nil {
			return
		}

		for range staff {
			if _, err := fmt.Scan(&operation, &degree); err != nil {
				return
			}

			switch operation {
			case ">=":
				temp.min = max(degree, temp.min)
			case "<=":
				temp.max = min(degree, temp.max)
			}

			if temp.min <= temp.max {
				fmt.Println(temp.min)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
