package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(x any) {
	if val, ok := x.(int); ok {
		*h = append(*h, val)
	}
}

func (h *MaxHeap) Pop() any {
	old := *h
	length := len(old)
	result := old[length-1]
	*h = old[:length-1]

	return result
}

func main() {
	var count int

	for {
		_, err := fmt.Scan(&count)
		if err != nil {
			break
		}

		maxHeap := &MaxHeap{}
		heap.Init(maxHeap)

		var value int

		for range count {
			_, _ = fmt.Scan(&value)
			heap.Push(maxHeap, value)
		}

		var targetK int

		_, _ = fmt.Scan(&targetK)

		if targetK > count || targetK < 1 {
			fmt.Println(-1)
		} else {
			for range targetK - 1 {
				heap.Pop(maxHeap)
			}

			fmt.Println(heap.Pop(maxHeap))
		}
	}
}
