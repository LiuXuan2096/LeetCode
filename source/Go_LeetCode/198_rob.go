package Go_LeetCode

// 给定一个数组arr，在不能取相邻数的情况下，返回所有组合中的最大累加和
// 思路：
// 定义dp[i] : 表示arr[0...i]范围上，在不能取相邻数的情况下，返回所有组合中的最大累加和
// 在arr[0...i]范围上，在不能取相邻数的情况下，得到的最大累加和，可能性分类：
// 可能性 1) 选出的组合，不包含arr[i]。那么dp[i] = dp[i-1]
// 比如，arr[0...i] = {3,4,-4}，最大累加和是不包含i位置数的时候
//
// 可能性 2) 选出的组合，只包含arr[i]。那么dp[i] = arr[i]
// 比如，arr[0...i] = {-3,-4,4}，最大累加和是只包含i位置数的时候
//
// 可能性 3) 选出的组合，包含arr[i], 且包含arr[0...i-2]范围上的累加和。那么dp[i] = arr[i] + dp[i-2]
// 比如，arr[0...i] = {3,1,4}，最大累加和是3和4组成的7，因为相邻不能选，所以i-1位置的数要跳过
//
// 综上所述：dp[i] = Max { dp[i-1], arr[i] , arr[i] + dp[i-2] }
func rob(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	N := len(nums)
	if N == 1 {
		return nums[0]
	}
	if N == 2 {
		return max_198(nums[0], nums[1])
	}
	dp := make([]int, N)
	dp[0] = nums[0]
	dp[1] = max_198(nums[0], nums[1])
	for i := 2; i < N; i++ {
		dp[i] = max_198(max_198(dp[i-1], nums[i]), nums[i]+dp[i-2])
	}
	return dp[N-1]
}

func max_198(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
