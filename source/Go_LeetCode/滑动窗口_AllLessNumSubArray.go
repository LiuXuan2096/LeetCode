package Go_LeetCode

import (
	"container/list"
)

// NumSubArray_滑动窗口
/*
给定一个整型数组arr，和一个整数num
某个arr中的子数组sub，如果想达标，必须满足：
sub中最大值 – sub中最小值 <= num，
返回arr中达标子数组的数量
*/
func NumSubArray_滑动窗口(arr []int, num int) int {
	if arr == nil || len(arr) == 0 || num < 0 {
		return 0
	}
	N := len(arr)
	count := 0
	maxWindow := list.New()
	minWindow := list.New()
	R := 0

	for L := 0; L < N; L++ {
		for R < N {
			for maxWindow.Len() != 0 && arr[maxWindow.Back().Value.(int)] <= arr[R] {
				maxWindow.Remove(maxWindow.Back())
			}
			maxWindow.PushBack(R)

			for minWindow.Len() != 0 && arr[minWindow.Back().Value.(int)] >= arr[R] {
				minWindow.Remove(minWindow.Back())
			}
			minWindow.PushBack(R)

			if arr[maxWindow.Front().Value.(int)]-arr[minWindow.Front().Value.(int)] > num {
				break
			} else {
				R++
			}
		}
		count += R - L

		if maxWindow.Front().Value.(int) == L {
			maxWindow.Remove(maxWindow.Front())
		}

		if minWindow.Front().Value.(int) == L {
			minWindow.Remove(minWindow.Front())
		}
	}

	return count
}
