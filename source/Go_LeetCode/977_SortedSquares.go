package Go_LeetCode

import "math"

// 时间复杂度O(N)，额外空间复杂度O(1)
func sortedSquares(nums []int) []int {
	N := len(nums)
	ans := make([]int, N)
	L := 0
	R := N - 1
	for i, k, j := L, R, R; i <= j; k-- {
		if pow2(nums[i]) > pow2(nums[j]) {
			ans[k] = pow2(nums[i])
			i++
		} else {
			ans[k] = pow2(nums[j])
			j--
		}
	}
	return ans
}

func pow2(a int) int {
	return int(math.Pow(float64(a), 2))
}
