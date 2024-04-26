package Go_LeetCode

import (
	"container/heap"
	"math"
)

/*
arr中求子数组的累加和是<=K的并且是最大的，返回这个最大的累加和
*/

// 实现一个小跟堆

type IntHeap_14_02 []int

func (h IntHeap_14_02) Len() int {
	return len(h)
}

func (h IntHeap_14_02) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap_14_02) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap_14_02) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap_14_02) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func GetMaxLessOrEqualK(arr []int, K int) int {
	// 记录i之前的，前缀和，按照有序表组织
	// 使用最小堆来模拟TreeSet的ceiling功能
	h := &IntHeap_14_02{}
	heap.Init(h)
	// 一个数也没有的时候，就已经有一个前缀和是0了
	heap.Push(h, 0)
	max := math.MinInt32
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i] // sum -> arr[0..i]
		// 使用最小堆来找到第一个大于或等于sum - K的元素
		for h.Len() > 0 && (*h)[0] <= sum-K {
			heap.Pop(h)
		}

		if h.Len() > 0 {
			max = max_32(max, sum-(*h)[0])
		}
		heap.Push(h, sum)
	}
	return max
}
