package Go_LeetCode

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	pre := nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		pre = maxInt(nums[i], nums[i]+pre)
		max = maxInt(pre, max)
	}
	return max
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
