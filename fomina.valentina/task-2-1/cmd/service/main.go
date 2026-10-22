package main

import (
	"container/heap"
	"fmt"
	"os"
)

type MyMaxHeap []int

func (h *MyMaxHeap) Len() int {
	return len(*h)
}

func (h *MyMaxHeap) Less(i int, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *MyMaxHeap) Swap(i int, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MyMaxHeap) Push(elem any) {
	value, ok := elem.(int)
	if !ok {
		fmt.Fprintln(os.Stderr, "Non-integer element cannot be pushed")

		return
	}

	*h = append(*h, value)
}

func (h *MyMaxHeap) Pop() any {
	length := len(*h)
	lastValue := (*h)[length-1]
	*h = (*h)[:length-1]

	return lastValue
}

func main() {
	var dishesCount int

	_, err := fmt.Scan(&dishesCount)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)

		return
	}

	heapOfDishes := &MyMaxHeap{}
	heap.Init(heapOfDishes)

	for range dishesCount {
		var curDish int

		_, err = fmt.Scan(&curDish)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)

			return
		}

		heap.Push(heapOfDishes, curDish)
	}

	var todayBest int

	_, err = fmt.Scan(&todayBest)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)

		return
	}

	for range todayBest - 1 {
		heap.Pop(heapOfDishes)
	}

	fmt.Println(heap.Pop(heapOfDishes))
}
