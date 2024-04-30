package Go_LeetCode

func firstMissingPositive(nums []int) int {
	// L是盯着的位置
	// 0 ~ L-1有效区
	L := 0
	R := len(nums)

	for L != R {
		// 有效区可以扩充的情况
		if nums[L] == L+1 {
			L++
		} else if nums[L] <= L || nums[L] > R || nums[L] == nums[nums[L]-1] {
			// 垃圾区可以扩充的三种情况
			// 1.nums[L] < L+1
			// 2.nums[L] > R 因为R代表当前可能得到的答案中的最大值，这个位置的值吧R还大
			// 说明这个位置的值肯定不会是答案 而且废掉了一个位置 R-1
			// 3.比如 nums[17] = 22 = nums[21] 这种情况说明有重复值 还是废掉了一个位置 R-1
			nums[L], nums[R-1] = nums[R-1], nums[L]
			R--
		} else {
			// L < 当前位置的数字 < R
			// 所以将该数字交换到它应该在的位置
			nums[L], nums[nums[L]-1] = nums[nums[L]-1], nums[L]
		}
	}
	return L + 1
}
