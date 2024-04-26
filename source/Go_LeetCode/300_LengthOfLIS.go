package Go_LeetCode

func lengthOfLIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	// ends数组
	// ends[i]表示 : 目前所有长度为i+1的递增子序列的最小结尾
	ends := make([]int, len(nums))
	// 根据含义, 一开始ends[0] = arr[0]
	ends[0] = nums[0]
	// ends有效区范围是0...right，right往右为无效区
	// 所以一开始right = 0, 表示有效区只有0...0范围
	right := 0
	// 最长递增子序列的长度
	// 全局变量，抓取每一步的答案，取最大的结果
	max := 1

	for i := 1; i < len(nums); i++ {
		l, r := 0, right
		for l <= r {
			m := (l + r) / 2
			// 在ends[l...r]范围上二分
			// 如果 当前数(arr[i]) > ends[m]，砍掉左侧
			// 如果 当前数(arr[i]) <= ends[m]，砍掉右侧
			// 整个二分就是在ends里寻找 >= 当前数(arr[i])的最左位置
			// 就是从while里面出来时，l所在的位置。
			// 如果ends中不存在 >= 当前数(arr[i])的情况，将返回有效区的越界位置
			// 也就是从while里面出来时，l所在的位置，是有效区的越界位置
			// 比如 : ends = { 3, 5, 9, 12, 再往右无效}
			// 如果当前数为8, 从while里面出来时，l将来到2位置
			// 比如 : ends = { 3, 5, 9, 12, 再往右无效}
			// 如果当前数为13, 从while里面出来时，l将来到有效区的越界位置，4位置
			if nums[i] > ends[m] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
		// 从while里面出来，看l的位置
		// 如果l比right大，说明扩充了有效区，那么right变量要随之变大
		// 如果l不比right大，说明l没有来到有效区的越界位置，right不变
		right = max_300(l, right)
		// l的位置，就是当前数应该填到ends数组里的位置
		ends[l] = nums[i]
		// 更新全局变量
		max = max_300(max, l+1)
	}
	return max
}

func max_300(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
