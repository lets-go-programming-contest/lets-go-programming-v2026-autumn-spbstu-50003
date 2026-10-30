package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	var n int
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}

		h := &MaxHeap{}
		heap.Init(h)

		for range n {
			var val int
			_, _ = fmt.Scan(&val)
			heap.Push(h, val)
		}

		var m int
		_, _ = fmt.Scan(&m)

		for range m {
			var k int
			_, _ = fmt.Scan(&k)

			if k > n || k < 1 {
				fmt.Println(-1)
			} else {
				temp := make([]int, 0, k)
				for range k {
					temp = append(temp, heap.Pop(h).(int))
				}
				result := temp[k-1]
				for _, v := range temp {
					heap.Push(h, v)
				}
				fmt.Println(result)
			}
		}
	}
}
