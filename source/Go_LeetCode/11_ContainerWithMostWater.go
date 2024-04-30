package Go_LeetCode

func maxArea(height []int) int {
	maxWater := 0
	l := 0
	r := len(height) - 1
	for l < r {
		maxWater = max(maxWater, min(height[l], height[r])*(r-l))
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return maxWater
}

//func max(a, b int) int {
//	if a >= b {
//		return a
//	}
//	return b
//}

//func min(a, b int) int {
//	if a <= b {
//		return a
//	}
//	return b
//}
