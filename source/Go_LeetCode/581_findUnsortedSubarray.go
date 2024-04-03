package Go_LeetCode

import "math"

func findUnsortedSubarray(nums []int) int {
	if nums == nil || len(nums) < 2 {
		return 0
	}
	N := len(nums)
	right := -1
	max := math.MinInt32
	for i := 0; i < N; i++ {
		if max > nums[i] {
			right = i
		}
		max = myMax(max, nums[i])
	}
	min := math.MaxInt32
	left := N
	for i := N - 1; i >= 0; i-- {
		if min < nums[i] {
			left = i
		}
		min = myMin(min, nums[i])
	}
	return myMax(0, right-left+1)
}

func myMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func myMin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
