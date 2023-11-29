package Go_LeetCode

import (
	"container/heap"
)

type program struct {
	capital int
	profit  int
}

type pHeap struct {
	programHeap []program
	flag        bool // true是小跟堆 false是大根堆
}

func (h *pHeap) Len() int {
	return len(h.programHeap)
}

func (h *pHeap) Less(i, j int) bool {
	if h.flag {
		// 小跟堆 根据成本比较
		return h.programHeap[i].capital < h.programHeap[j].capital
	} else {
		// 大根堆 根据利润比较
		return h.programHeap[i].profit > h.programHeap[j].profit
	}
}

func (h *pHeap) Swap(i, j int) {
	h.programHeap[i], h.programHeap[j] = h.programHeap[j], h.programHeap[i]
}

func (h *pHeap) Top() int {
	if h.flag {
		// 小跟堆的top返回 成本
		return h.programHeap[0].capital
	} else {
		// 大根堆的top返回 利润
		return h.programHeap[0].profit
	}
}

func (h *pHeap) Push(x interface{}) {
	h.programHeap = append(h.programHeap, x.(program))
}

func (h *pHeap) Pop() interface{} {
	old := h
	length := len(h.programHeap)
	x := h.programHeap[length-1]
	h.programHeap = old.programHeap[0 : length-1]
	return x
}

func FindMaximizedCapital(k int, w int, profits []int, capital []int) int {
	var minHeap pHeap = pHeap{flag: true}
	heap.Init(&minHeap)
	var maxHeap = pHeap{flag: false}
	heap.Init(&maxHeap)

	for i := 0; i < len(profits); i++ {
		heap.Push(&minHeap, program{capital[i], profits[i]})
	}
	for i := 0; i < k; i++ {
		for minHeap.Len() > 0 && minHeap.Top() <= w {
			heap.Push(&maxHeap, heap.Pop(&minHeap))
		}
		if maxHeap.Len() == 0 {
			return w
		}
		w += maxHeap.Top()
		heap.Pop(&maxHeap)
	}
	return w
}
