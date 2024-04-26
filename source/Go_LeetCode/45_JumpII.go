package Go_LeetCode

func jump(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	step := 0 //当前走的步数
	cur := 0  //当前步数下能走到的最远距离
	next := 0 //再多走一步 能走到的最远距离
	for i := 0; i < len(nums); i++ {
		if cur < i {
			step++
			cur = next
		}
		next = max_45(next, i+nums[i])
	}
	return step
}

func max_45(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
