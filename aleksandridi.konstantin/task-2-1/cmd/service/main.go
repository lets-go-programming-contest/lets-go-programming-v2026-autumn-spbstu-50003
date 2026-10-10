package main

import "fmt"

type Conditioner struct {
	maxTemperature int
	minTemperature int
}

func NewConditioner() Conditioner {
	return Conditioner{
		maxTemperature: 30,
		minTemperature: 15,
	}
}

func (c *Conditioner) SetTemperature(operation string, value int) int {
	if c.minTemperature <= c.maxTemperature {
		switch operation {
		case "<=":
			if c.maxTemperature > value {
				c.maxTemperature = value
			}
		case ">=":
			if c.minTemperature < value {
				c.minTemperature = value
			}
		}
	}
	if c.minTemperature <= c.maxTemperature {
		return c.minTemperature
	}
	return -1
}

func main() {
	var departmentCount int
	if _, err := fmt.Scan(&departmentCount); err != nil {
		return
	}

	for i := 0; i < departmentCount; i++ {
		var staffCount int
		if _, err := fmt.Scan(&staffCount); err != nil {
			return
		}

		conditioner := NewConditioner()

		for j := 0; j < staffCount; j++ {
			var (
				operation string
				value     int
			)
			if _, err := fmt.Scan(&operation, &value); err != nil {
				return
			}
			fmt.Println(conditioner.SetTemperature(operation, value))
		}
	}
}
