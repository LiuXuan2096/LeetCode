package Go_LeetCode

import (
	"container/list"
	"github.com/gammazero/deque"
)

func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k < 1 || len(nums) < k {
		return nil
	}
	// qmax 窗口最大值的更新结构
	// 存放下标
	var q deque.Deque[int]
	var ans = make([]int, len(nums)-k+1)
	var index = 0

	for R := 0; R < len(nums); R++ {
		if q.Len() != 0 && nums[q.Back()] <= nums[R] {
			q.PopBack()
		}
		q.PushBack(R)
		// fmt.Println("每次加入窗口的值: ", nums[R])
		if q.Front() == R-k {
			q.PopFront()
		}
		if R >= k-1 {
			ans[index] = nums[q.Front()]
			index++
		}
	}
	return ans
}

func MgetMaxWindow(arr []int, w int) []int {
	if arr == nil || w < 1 || len(arr) < w {
		return nil
	}

	var qmax list.List
	res := make([]int, len(arr)-w+1)
	index := 0

	for R := 0; R < len(arr); R++ {
		for qmax.Len() > 0 && arr[qmax.Back().Value.(int)] <= arr[R] {
			qmax.Remove(qmax.Back())
		}
		qmax.PushBack(R)

		if qmax.Front().Value.(int) == R-w {
			qmax.Remove(qmax.Front())
		}

		if R >= w-1 {
			res[index] = arr[qmax.Front().Value.(int)]
			index++
		}
	}

	return res
}
