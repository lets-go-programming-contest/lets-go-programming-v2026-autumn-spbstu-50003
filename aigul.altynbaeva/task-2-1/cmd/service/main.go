package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	
	// Читаем все тесты до конца файла (EOF)
	for fmt.Scan(&n) == 1 {
		a := make([]int, n)
		for i := range n {
			fmt.Scan(&a[i])
		}

		sort.Sort(sort.Reverse(sort.IntSlice(a)))

		var k int
		fmt.Scan(&k)

		if k > n || k < 1 {
			fmt.Println(-1)
		} else {
			fmt.Println(a[k-1])
		}
	}
}
