package Go_LeetCode

import "math"

func LargestRectangleArea(heights []int) int {
	if heights == nil || len(heights) == 0 {
		return 0
	}
	var N = len(heights)
	var stack = make([]int, N)
	var stackIndex = -1
	var maxArea = 0
	for i := 0; i < N; i++ {
		for stackIndex != -1 && heights[i] <= heights[stack[stackIndex]] {
			var j = stack[stackIndex]
			stackIndex--
			var k int
			if stackIndex == -1 {
				k = -1
			} else {
				k = stack[stackIndex]
			}
			var curArea = (i - k - 1) * heights[j]
			maxArea = int(math.Max(float64(maxArea), float64(curArea)))
		}
		stack[stackIndex+1] = i
		stackIndex++
	}
	for stackIndex != -1 {
		var j = stack[stackIndex]
		stackIndex--
		var k int
		if stackIndex == -1 {
			k = -1
		} else {
			k = stack[stackIndex]
		}
		var curArea = (N - k - 1) * heights[j]
		maxArea = int(math.Max(float64(maxArea), float64(curArea)))

	}
	return maxArea
}
