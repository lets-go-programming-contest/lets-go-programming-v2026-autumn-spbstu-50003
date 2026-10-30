package main

import (
	"fmt"
	"sort"
)

func main() {
	var elementsCount int

	for {
		_, err := fmt.Scan(&elementsCount)
		if err != nil {
			break
		}

		elements := make([]int, elementsCount)

		for i := range elementsCount {
			_, _ = fmt.Scan(&elements[i])
		}


		sort.Sort(sort.Reverse(sort.IntSlice(elements)))

		var queriesCount int
		_, err = fmt.Scan(&queriesCount)
		if err != nil {
			break
		}


		for range queriesCount {
			var kth int
			_, _ = fmt.Scan(&kth)


			if kth > elementsCount || kth < 1 {
				fmt.Println(-1)
			} else {

				fmt.Println(elements[kth-1])
			}
		}
	}
}
