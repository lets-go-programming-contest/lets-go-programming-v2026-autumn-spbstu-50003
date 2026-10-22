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

func (r *Range) Apply(mod string, temp int) int {
	if mod == opLessOrEqual && temp < r.Max {
		r.Max = temp
	} else if mod == opGreaterOrEqual && temp > r.Min {
		r.Min = temp
	}

	if r.Min <= r.Max {
		return r.Min
	}

	return -1
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

			fmt.Println(rng.Apply(mod, temp))
		}
	}
}
