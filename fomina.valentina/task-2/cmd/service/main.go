package main

import (
	"container/heap"
	"fmt"
	"os"
)

type DishMaxHeap []int

func (h *DishMaxHeap) Len() int {
	return len(*h)
}

func (h *DishMaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *DishMaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *DishMaxHeap) Push(item any) {
	dishRating, ok := item.(int)
	if !ok {
		fmt.Fprintln(os.Stderr, "Ошибка: можно добавлять только целые числа")
		return
	}
	*h = append(*h, dishRating)
}

func (h *DishMaxHeap) Pop() any {
	oldLength := len(*h)
	lastElement := (*h)[oldLength-1]
	*h = (*h)[:oldLength-1]
	return lastElement
}

func main() {
	var dishesCount int
	if _, err := fmt.Scan(&dishesCount); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения количества блюд:", err)
		return
	}

	ratingHeap := &DishMaxHeap{}
	heap.Init(ratingHeap)

	for i := 0; i < dishesCount; i++ {
		var dishRating int
		if _, err := fmt.Scan(&dishRating); err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка чтения рейтинга блюда:", err)
			return
		}
		heap.Push(ratingHeap, dishRating)
	}

	var dishPosition int
	if _, err := fmt.Scan(&dishPosition); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка чтения позиции блюда:", err)
		return
	}

	for i := 0; i < dishPosition-1; i++ {
		heap.Pop(ratingHeap)
	}

	fmt.Println(heap.Pop(ratingHeap))
}
