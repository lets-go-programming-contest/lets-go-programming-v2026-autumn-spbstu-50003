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

func (t *TempRange) Apply(operator string, temp int) int {
	if t.Min > t.Max {
		t.Reset()
	}

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

	if t.Min > t.Max {
		return -1
	}
	
	return t.Max
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

			if _, err := fmt.Scan(&operator, &temp); err != nil {
				return
			}

			fmt.Println(currentRange.Apply(operator, temp))
		}
	}
}
