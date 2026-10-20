package main

import (
	"fmt"
)

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	if n < 1 || n > 1000 {
		return
	}

	for dept := 0; dept < n; dept++ {
		var k int
		if _, err := fmt.Scan(&k); err != nil {
			break
		}

		if k < 1 || k > 1000 {
			break
		}

		minTemp := 15
		maxTemp := 30

		for emp := 0; emp < k; emp++ {
			var (
				op  string
				val int
			)
			if _, err := fmt.Scan(&op, &val); err != nil {
				break
			}

			if op == ">=" {
				if val > minTemp {
					minTemp = val
				}
			} else if op == "<=" {
				if val < maxTemp {
					maxTemp = val
				}
			}

			if minTemp <= maxTemp {
				fmt.Println(minTemp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
