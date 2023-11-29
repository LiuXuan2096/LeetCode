package Go_LeetCode

type IntHeap struct {
	heap []int
	bool // true是小跟堆 false是大根堆
}

func (h *IntHeap) Len() int {
	return len(h.heap)
}

func (h *IntHeap) Less(i, j int) bool {
	if h.bool {
		return h.heap[i] < h.heap[j]
	} else {
		return h.heap[i] > h.heap[j]
	}
}

func (h *IntHeap) Swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func (h *IntHeap) Push(x interface{}) {
	h.heap = append(h.heap, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := h
	length := len(old.heap)
	top := old.heap[length-1]
	h.heap = old.heap[0 : length-1]
	return top
}

func (h *IntHeap) Top() int {
	return h.heap[0]
}
