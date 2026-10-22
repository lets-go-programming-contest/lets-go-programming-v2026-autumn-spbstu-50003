package main

import "fmt"

const (
	minTemp = 15
	maxTemp = 30

	opLessOrEqual    = "<="
	opGreaterOrEqual = ">="
)

type Range struct {
	Min int
	Max int
}

func main() {
	var sites int

	if _, err := fmt.Scan(&sites); err != nil {
		return
	}

	for range sites {
		rng := Range{Min: minTemp, Max: maxTemp}

		var siteEmployees int

		if _, err := fmt.Scan(&siteEmployees); err != nil {
			return
		}

		for range siteEmployees {
			var (
				mod  string
				temp int
			)

			if _, err := fmt.Scan(&mod, &temp); err != nil {
				return
			}

			if mod == "<=" {
				if temp < rng.Max {
					rng.Max = temp
				}
			} else if mod == ">=" {
				if temp > rng.Min {
					rng.Min = temp
				}
			}

			if rng.Min <= rng.Max {
				fmt.Printf("%d\n", rng.Min)
			} else {
				fmt.Printf("-1\n")
			}
		}
	}
}
