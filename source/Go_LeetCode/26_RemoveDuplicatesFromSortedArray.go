package Go_LeetCode

func removeDuplicates_26(nums []int) int {
	if nums == nil {
		return 0
	}
	if len(nums) < 2 {
		return len(nums)
	}
	done := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[done] {
			done++
			nums[done] = nums[i]
		}
	}
	return done + 1
}
