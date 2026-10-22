package main

import "fmt"

const (
	lowTemp  = 15
	highTemp = 30
)

type temperatureLimits struct {
	minBound int
	maxBound int
}

func (t *temperatureLimits) updateRange(sign string, tempValue int) int {
	switch sign {
	case ">=":
		if tempValue > t.minBound {
			t.minBound = tempValue
		}
	case "<=":
		if tempValue < t.maxBound {
			t.maxBound = tempValue
		}
	}

	if t.minBound <= t.maxBound {
		return t.minBound
	}

	return -1
}

func main() {
	var deptCount int
	if _, err := fmt.Scan(&deptCount); err != nil {
		fmt.Println(err)

		return
	}

	for i := 0; i < deptCount; i++ {
		var staffCount int
		if _, err := fmt.Scan(&staffCount); err != nil {
			fmt.Println(err)

			return
		}

		limits := temperatureLimits{minBound: lowTemp, maxBound: highTemp}

		for j := 0; j < staffCount; j++ {
			var (
				condition string
				temp      int
			)

			if _, err := fmt.Scan(&condition, &temp); err != nil {
				fmt.Println(err)

				return
			}

			fmt.Println(limits.updateRange(condition, temp))
		}
	}
}
