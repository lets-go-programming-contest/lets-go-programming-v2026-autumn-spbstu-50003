package main

import "fmt"

const (
	minTemp = 15
	maxTemp = 30
)

type tempRange struct {
	low, high int
}

func (r tempRange) update(condition string, limit int) tempRange {
	if condition == ">=" && limit > r.low {
		r.low = limit
	}

	if condition == "<=" && limit < r.high {
		r.high = limit
	}

	return r
}

func (r tempRange) result() int {
	if r.low > r.high {
		return -1
	}

	return r.low
}

func main() {
	var numDepts int

	if _, err := fmt.Scan(&numDepts); err != nil {
		fmt.Println("invalid number of departments", err)

		return
	}

	for range numDepts {
		var numEmps int

		if _, err := fmt.Scan(&numEmps); err != nil {
			fmt.Println("invalid number of employees", err)

			return
		}

		temps := tempRange{low: minTemp, high: maxTemp}

		for range numEmps {
			var (
				condition string
				limit     int
			)

			if _, err := fmt.Scan(&condition, &limit); err != nil {
				fmt.Println("invalid condition", err)

				return
			}

			temps = temps.update(condition, limit)

			fmt.Println(temps.result())
		}
	}
}
