package main

import "fmt"

type Range struct{ Min, Max int }

func main() {
	var (
		sites         int
		siteEmployees int
		mod           string
		temp          int
		rng           Range
	)

	if _, err := fmt.Scan(&sites); err != nil {
		return
	}

	for i := 0; i < sites; i++ {
		rng = Range{15, 30}
		if _, err := fmt.Scan(&siteEmployees); err != nil {
			return
		}
		for k := 0; k < siteEmployees; k++ {
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
