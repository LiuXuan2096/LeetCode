package Go_LeetCode

import "container/list"

/*
给定一个整型数组arr，和一个整数num
某个arr中的子数组sub，如果想达标，必须满足：sub中最大值 – sub中最小值 <= num，
返回arr中达标子数组的数量
*/

func Num(arr []int, sum int) int {
	if arr == nil || len(arr) == 0 || sum < 0 {
		return 0
	}
	var N = len(arr)
	var count = 0
	var maxWindow list.List
	var minWindow list.List
	var R = 0
	for L := 0; L < N; L++ {
		for R < N {
			for maxWindow.Len() > 0 && arr[maxWindow.Back().Value.(int)] <= arr[R] {
				maxWindow.Remove(maxWindow.Back())
			}
			maxWindow.PushBack(R)
			for minWindow.Len() > 0 && arr[minWindow.Back().Value.(int)] >= arr[R] {
				minWindow.Remove(maxWindow.Back())
			}
			minWindow.PushBack(R)
			if arr[maxWindow.Front().Value.(int)]-arr[minWindow.Front().Value.(int)] > sum {
				break // 此时R的值是第一次不达标的位置
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
