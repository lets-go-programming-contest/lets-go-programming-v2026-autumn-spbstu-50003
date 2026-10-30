package main

import "fmt"

const (
	MinAllowedTemp = 15
	MaxAllowedTemp = 30
)

type TempRange struct {
	Min int
	Max int
}

func (t *TempRange) Reset() {
	t.Min = MinAllowedTemp
	t.Max = MaxAllowedTemp
}

func (t *TempRange) Apply(operator string, temp int) {
	switch operator {
	case ">=":
		if temp > t.Min {
			t.Min = temp
		}
	case "<=":
		if temp < t.Max {
			t.Max = temp
		}
	}
}

func (t *TempRange) GetResult() int {
	if t.Min > t.Max {
		return -1
	}

	return t.Min
}

func main() {
	var departments int

	for {
		_, err := fmt.Scan(&departments)
		if err != nil {
			break
		}

		var currentRange TempRange

		currentRange.Reset()

		for range departments {
			var employees int

			_, err := fmt.Scan(&employees)
			if err != nil {
				return
			}

			for range employees {
				var (
					operator string
					temp     int
				)

				_, err := fmt.Scan(&operator, &temp)
				if err != nil {
					return
				}

				currentRange.Apply(operator, temp)

				fmt.Println(currentRange.GetResult())
			}
		}
	}
}
