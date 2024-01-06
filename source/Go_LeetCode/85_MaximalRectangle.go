package Go_LeetCode

import (
	"container/list"
	"math"
)

func MaximalRectangle(matrix [][]byte) int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	var maxArea = 0
	var heights = make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				heights[j] = 0
			} else {
				heights[j] += 1
			}
		}
		maxArea = int(math.Max(float64(maxArea), float64(maxRecFromBottom(heights))))
	}
	return maxArea
}

func maxRecFromBottom(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}
	var maxArea = 0
	var stack list.List
	for i := 0; i < len(height); i++ {
		for stack.Len() > 0 && height[stack.Front().Value.(int)] >= height[i] {
			var j = stack.Front().Value.(int)
			stack.Remove(stack.Front())
			var k = -1
			if stack.Len() > 0 {
				k = stack.Front().Value.(int)
			}

			var curArea = (i - k - 1) * height[j]
			maxArea = int(math.Max(float64(maxArea), float64(curArea)))
		}
		stack.PushFront(i)
	}
	for stack.Len() > 0 {
		var j = stack.Front().Value.(int)
		stack.Remove(stack.Front())
		var k int
		if stack.Len() == 0 {
			k = -1
		} else {
			k = stack.Front().Value.(int)
		}
		var curArea = (len(height) - k - 1) * height[j]
		maxArea = int(math.Max(float64(maxArea), float64(curArea)))
	}
	return maxArea
}

func MmaximalRectangle(mapData [][]byte) int {
	if mapData == nil || len(mapData) == 0 || len(mapData[0]) == 0 {
		return 0
	}

	maxArea := 0
	height := make([]int, len(mapData[0]))

	for i := 0; i < len(mapData); i++ {
		for j := 0; j < len(mapData[0]); j++ {
			if mapData[i][j] == '0' {
				height[j] = 0
			} else {
				height[j] += 1
			}
		}
		maxArea = int(math.Max(float64(MmaxRecFromBottom(height)), float64(maxArea)))
	}

	return maxArea
}

func MmaxRecFromBottom(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}

	maxArea := 0
	stack := make([]int, 0)

	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && height[i] <= height[stack[len(stack)-1]] {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			k := -1
			if len(stack) > 0 {
				k = stack[len(stack)-1]
			}
			curArea := (i - k - 1) * height[j]
			maxArea = int(math.Max(float64(maxArea), float64(curArea)))
		}
		stack = append(stack, i)
	}

	for len(stack) > 0 {
		j := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k := -1
		if len(stack) > 0 {
			k = stack[len(stack)-1]
		}
		curArea := (len(height) - k - 1) * height[j]
		maxArea = int(math.Max(float64(maxArea), float64(curArea)))
	}

	return maxArea
}
