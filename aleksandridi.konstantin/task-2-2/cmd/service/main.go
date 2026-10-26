package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x any) {
	value, isTrue := x.(int)
	if !isTrue {
		fmt.Println("Not the right type!")

		return
	}

	*h = append(*h, value)
}

func (h *IntHeap) Pop() any {
	old := *h
	length := len(old)
	x := old[length-1]
	*h = old[0 : length-1]

	return x
}

func main() {
	var foodCount int

	if _, err := fmt.Scan(&foodCount); err != nil {
		fmt.Println(err)

		return
	}

	foodRating := make([]int, foodCount)

	for i := range foodCount {
		if _, err := fmt.Scan(&foodRating[i]); err != nil {
			fmt.Println(err)

			return
		}
	}

	var preference int

	if _, err := fmt.Scan(&preference); err != nil {
		fmt.Println(err)

		return
	}

	intHeap := &IntHeap{}
	heap.Init(intHeap)

	for _, value := range foodRating {
		heap.Push(intHeap, value)

		if intHeap.Len() > preference {
			heap.Pop(intHeap)
		}
	}

	fmt.Println((*intHeap)[0])
}
