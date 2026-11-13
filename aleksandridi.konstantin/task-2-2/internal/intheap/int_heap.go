package intheap

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	if i < 0 || i >= len(*h) || j < 0 || j >= len(*h) {
		return false
	}

	return (*h)[i] < (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	if i < 0 || i >= len(*h) || j < 0 || j >= len(*h) {
		return
	}

	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x any) {
	value, ok := x.(int)
	if !ok {
		return
	}

	*h = append(*h, value)
}

func (h *IntHeap) Pop() any {
	old := *h
	length := len(old)
	if length == 0 {
		return nil
	}

	x := old[length-1]
	*h = old[0 : length-1]

	return x
}
