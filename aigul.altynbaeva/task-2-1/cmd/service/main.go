package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	for fmt.Scan(&n) == 1 {
		a := make([]int, n)
		for i := range n {
			_, _ = fmt.Scan(&a[i])
		}
		sort.Sort(sort.Reverse(sort.IntSlice(a)))

		var m int
		_, _ = fmt.Scan(&m)
		for range m {
			var k int
			_, _ = fmt.Scan(&k)
			if k < 1 || k > n {
				fmt.Println(-1)
			} else {
				fmt.Println(a[k-1])
			}
		}
	}
}
