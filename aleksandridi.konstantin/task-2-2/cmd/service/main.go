package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {

	return len(h)
}

func (h IntHeap) Less(i, j int) bool {

	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println(err)

		return
	}
	foodRaiting := make([]int, n)
	for i := range n {
		fmt.Scan(&foodRaiting[i])
	}
	var k int
	if _, err := fmt.Scan(&k); err != nil {
		fmt.Println(err)

		return
	}
	h := &IntHeap{}
	heap.Init(h)

	for _, value := range foodRaiting {
		heap.Push(h, value)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	fmt.Println((*h)[0])
}
