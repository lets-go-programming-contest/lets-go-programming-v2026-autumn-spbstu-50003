package main

import "fmt"

const (
	minLimit = 15
	maxLimit = 30
)

type temperatureRange struct {
	min int
	max int
}

func (r *temperatureRange) apply(mod string, temp int) int {
	switch mod {
	case ">=":
		if temp > r.min {
			r.min = temp
		}
	case "<=":
		if temp < r.max {
			r.max = temp
		}
	}

	if r.min <= r.max {
		return r.min
	}

	return -1
}

func main() {
	var departments int
	if _, err := fmt.Scan(&departments); err != nil {
		fmt.Println(err)

		return
	}

	for range departments {
		var employees int
		if _, err := fmt.Scan(&employees); err != nil {
			fmt.Println(err)

			return
		}

		rng := temperatureRange{min: minLimit, max: maxLimit}

		for range employees {
			var (
				mod  string
				temp int
			)

			if _, err := fmt.Scan(&mod, &temp); err != nil {
				fmt.Println(err)

				return
			}

			fmt.Println(rng.apply(mod, temp))
		}
	}
}
