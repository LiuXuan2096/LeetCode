package Go_LeetCode

//
//import (
//	"container/heap"
//	"math"
//)
//
//var (
//	next  []int
//	numsC [][]int
//)
//
//func smallestRange(nums [][]int) []int {
//	numsC = nums
//	rangeLeft, rangeRight := 0, math.MaxInt32
//	minRange := rangeRight - rangeLeft
//	maxV := math.MinInt32
//	size := len(nums)
//	next = make([]int, size) //next[i]的含义是当前要比较的i号链表的值的索引
//	h := &IHeap{}
//	heap.Init(h)
//
//	for i := 0; i < size; i++ {
//		heap.Push(h, i) // 堆里放的是每个数组的索引
//		maxV = max(maxV, nums[i][0])
//	}
//
//	for {
//		minIndex := heap.Pop(h).(int) //当前的最小值在几号链表
//		curRange := maxV - nums[minIndex][next[minIndex]]
//		if curRange < minRange {
//			minRange = curRange
//			rangeLeft, rangeRight = nums[minIndex][next[minIndex]], maxV
//		}
//		next[minIndex]++ //minIndex号链表的下一个最小值的索引+1
//		if next[minIndex] == len(nums[minIndex]) {
//			break //说明有一个链表遍历完了，循环结束
//		}
//		heap.Push(h, minIndex)
//		maxV = max(maxV, nums[minIndex][next[minIndex]])
//	}
//
//	return []int{rangeLeft, rangeRight}
//}
//
//type IHeap []int
//
//func (h IHeap) Len() int {
//	return len(h)
//}
//
//func (h IHeap) Less(i, j int) bool {
//	return numsC[h[i]][next[h[i]]] < numsC[h[j]][next[h[j]]]
//}
//
//func (h IHeap) Swap(i, j int) {
//	h[i], h[j] = h[j], h[i]
//}
//
//func (h *IHeap) Push(x interface{}) {
//	*h = append(*h, x.(int))
//}
//
//func (h *IHeap) Pop() interface{} {
//	old := *h
//	n := len(old)
//	x := old[n-1]
//	*h = old[0 : n-1]
//	return x
//}
