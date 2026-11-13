package main

import (
	"container/heap"
	"fmt"

	"github.com/neBr3aD/task-2-2/internal/intheap"
)

func main() {
	var foodCount int
	if _, err := fmt.Scan(&foodCount); err != nil {
		fmt.Println(fmt.Errorf("failed to read food count: %w", err))
		return
	}

	foodRating := make([]int, foodCount)
	for i := range foodCount {
		if _, err := fmt.Scan(&foodRating[i]); err != nil {
			fmt.Println(fmt.Errorf("failed to read rating for food #%d: %w", i+1, err))
			return
		}
	}

	var preference int
	if _, err := fmt.Scan(&preference); err != nil {
		fmt.Println(fmt.Errorf("failed to read preference: %w", err))
		return
	}

	h := &intheap.IntHeap{}
	heap.Init(h)

	for _, value := range foodRating {
		heap.Push(h, value)

		if h.Len() > preference {
			heap.Pop(h)
		}
	}

	if h.Len() > 0 {
		fmt.Println((*h)[0])
	}
}
