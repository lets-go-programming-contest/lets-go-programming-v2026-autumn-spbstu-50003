package main

import "fmt"

const (
	LowestTemp  = 15
	HighestTemp = 30
)

type Border struct {
	Min int
	Max int
}

func main() {
	var blocks int

	if _, err := fmt.Scan(&blocks); err != nil {
		fmt.Println("failed to read blocks:", err)

		return
	}

	for range blocks {
		var staffCount int

		if _, err := fmt.Scan(&staffCount); err != nil {
			fmt.Println("failed to read staff count")

			return
		}

		bounds := Border{
			Min: LowestTemp,
			Max: HighestTemp,
		}

		for range staffCount {
			var (
				sign string
				val  int
			)

			if _, err := fmt.Scan(&sign, &val); err != nil {
				fmt.Println("failed to read condition")

				return
			}

			switch sign {
			case ">=":
				if val > bounds.Min {
					bounds.Min = val
				}
			case "<=":
				if val < bounds.Max {
					bounds.Max = val
				}
			default:
				fmt.Println("unknown operator")

				return
			}

			if bounds.Min > bounds.Max {
				fmt.Println(-1)
			} else {
				fmt.Println(bounds.Min)
			}
		}
	}
}
