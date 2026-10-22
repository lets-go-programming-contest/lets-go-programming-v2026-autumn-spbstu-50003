package main

import "fmt"

const (
	minT = 15
	maxT = 30

	opNoLessThen = ">="
	opNoMoreThan = "<="
)

type Range struct {
	Low  int
	High int
}

func (r *Range) Update(sign string, value int) int {
	switch sign {
	case opNoMoreThan:
		if value < r.High {
			r.High = value
		}
	case opNoLessThen:
		if value > r.Low {
			r.Low = value
		}
	}

	if r.Low > r.High {
		return -1
	}

	return r.Low
}

func main() {
	var deptCount int

	if _, err := fmt.Scan(&deptCount); err != nil {
		fmt.Println("Failed to read number of departments:", err)
		return
	}
	for range deptCount {
		var staffCount int

		if _, err := fmt.Scan(&staffCount); err != nil {
			fmt.Println("Failed to read number of employees:", err)
			return
		}

		rng := Range{Low: minT, High: maxT}

		for range staffCount {
			var (
				sign  string
				value int
			)

			if _, err := fmt.Scan(&sign, &value); err != nil {
				fmt.Println("Failed to read condition:", err)
				return
			}

			fmt.Println(rng.Update(sign, value))
		}
	}
}
