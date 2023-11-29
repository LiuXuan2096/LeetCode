package Go_LeetCode

/*

这个程序有错误，100个用例只能通过18个
*/

import (
	"container/heap"
	"sort"
)

type intHeap []any

func minCost(n int, cuts []int) int {
	sort.Ints(cuts)

	lengthArray := make([]int, len(cuts)+1)
	lengthArray[0] = cuts[0]
	for i := 1; i < len(cuts); i++ {
		lengthArray[i] = cuts[i] - cuts[i-1]
	}
	lengthArray[len(cuts)] = n - cuts[len(cuts)-1]

	// 初始化小跟堆

	minHeap := &intHeap{}
	heap.Init(minHeap)
	for _, val := range lengthArray {
		heap.Push(minHeap, val)
	}
	var sum = 0
	var cur = 0
	for len(*minHeap) > 1 {
		cur = heap.Pop(minHeap).(int) + heap.Pop(minHeap).(int)
		sum += cur
		heap.Push(minHeap, cur)
	}
	return sum

}

func (h *intHeap) Push(x any) {
	// Push 和 Pop 使用 pointer receiver 作为参数
	// 因为它们不仅会对切片的内容进行调整，还会修改切片的长度。
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() any {
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

func (h *intHeap) Len() int {
	return len(*h)
}

func (h *intHeap) Less(i, j int) bool {
	// 如果实现大顶堆，则需要调整为大于号
	return (*h)[i].(int) < (*h)[j].(int)
}

func (h *intHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *intHeap) Top() any {
	return (*h)[0]
}
