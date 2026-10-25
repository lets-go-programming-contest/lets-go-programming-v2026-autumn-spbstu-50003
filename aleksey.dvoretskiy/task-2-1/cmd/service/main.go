package main

import "fmt"

func main() {
	var departments int
	if _, err := fmt.Scan(&departments); err != nil {
		return
	}

	for range departments {
		var employees int
		if _, err := fmt.Scan(&employees); err != nil {
			return
		}

		minTemp := 15
		maxTemp := 30

		for range employees {
			var (
				operator    string
				temperature int
			)

			if _, err := fmt.Scan(&operator, &temperature); err != nil {
				return
			}

			switch operator {
			case ">=":
				if temperature > minTemp {
					minTemp = temperature
				}
			case "<=":
				if temperature < maxTemp {
					maxTemp = temperature
				}
			}

			if minTemp > maxTemp {
				fmt.Println(-1)
			} else {
				fmt.Println(minTemp)
			}
		}
	}
}
