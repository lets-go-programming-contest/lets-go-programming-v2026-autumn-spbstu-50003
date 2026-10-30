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
		return
	}

	for range departments {
		var employees int
		if _, err := fmt.Scan(&employees); err != nil {
			return
		}

		currentRange := TempRange{
			Min: MinAllowedTemp,
			Max: MaxAllowedTemp,
		}

		for range employees {
			var operator string
			var temp int

			_, err := fmt.Scan(&operator, &temp)
			if err != nil {
				return
			}

			switch operator {
			case ">=":
				if temp > currentRange.Min {
					currentRange.Min = temp
				}
			case "<=":
				if temp < currentRange.Max {
					currentRange.Max = temp
				}
			}

			if currentRange.Min > currentRange.Max {
				fmt.Println(-1)
			} else {
				fmt.Println(currentRange.Min)
			}
		}
	}
}
