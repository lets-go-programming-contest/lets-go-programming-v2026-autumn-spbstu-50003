package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		n, err := strconv.Atoi(parts[0])
		if err != nil || n == 0 {
			continue
		}

		var currentRange TempRange
		currentRange.Reset()

		for i := 1; i <= n && i < len(parts); i++ {
			if i+1 < len(parts) {
				operator := parts[i]
				temp, err := strconv.Atoi(parts[i+1])
				if err == nil {
					currentRange.Apply(operator, temp)
				}
				i++
			}
		}

		fmt.Println(currentRange.GetResult())
	}
}
