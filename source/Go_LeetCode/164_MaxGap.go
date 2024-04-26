package Go_LeetCode

import "math"

func maximumGap(nums []int) int {
	if nums == nil || len(nums) < 2 {
		return 0
	}
	length := len(nums)
	min := math.MaxInt32
	max := math.MinInt32
	for i := 0; i < length; i++ {
		min = min_164(min, nums[i])
		max = max_164(max, nums[i])
	}
	if min == max {
		return 0
	}
	hasNum := make([]bool, length+1)
	maxs := make([]int, length+1)
	mins := make([]int, length+1)
	bid := 0
	for i := 0; i < length; i++ {
		bid = bucket(nums[i], length, min, max)
		if hasNum[bid] {
			mins[bid] = min_164(mins[bid], nums[i])
			maxs[bid] = max_164(maxs[bid], nums[i])
		} else {
			mins[bid] = nums[i]
			maxs[bid] = nums[i]
			hasNum[bid] = true
		}
	}
	res := 0
	lastMax := maxs[0]
	i := 1
	for i <= length {
		if hasNum[i] {
			res = max_164(res, mins[i]-lastMax)
			lastMax = maxs[i]
		}
		i++
	}
	return res
}

func min_164(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max_164(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func bucket(num, length, min, max int) int {
	// 一个桶的范围
	rangeVal := float64(max-min) / float64(length)
	// num和min之间的距离
	distance := float64(num - min)
	// 返回桶的编号，向下取整
	return int(math.Floor(distance / rangeVal))
}
