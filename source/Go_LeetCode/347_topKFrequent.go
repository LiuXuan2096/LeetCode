package Go_LeetCode

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	map_num := map[int]int{}
	for _, item := range nums {
		map_num[item]++
	}
	h := &IHeap{}
	heap.Init(h)
	// 所有元素入堆，堆的长度为k
	for key, value := range map_num {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([]int, k)
	// 按顺序返回堆中的元素
	for i := 0; i < k; i++ {
		res[k-1-i] = heap.Pop(h).([2]int)[0]
	}
	return res
}

// 构建小跟堆
type IHeap [][2]int

// Len Swap Less Push Pop
func (h IHeap) Len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
