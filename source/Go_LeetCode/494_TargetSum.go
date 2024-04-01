package Go_LeetCode

func findTargetSumWays_1(nums []int, target int) int {
	return process_494(nums, 0, target)
}

func process_494(arr []int, i, rest int) int {
	if i == len(arr) {
		if rest == 0 {
			return 1
		} else {
			return 0
		}
	}
	return process_494(arr, i+1, rest-arr[i]) + process_494(arr, i+1, rest+arr[i])
}

func findTargetSumWays(arr []int, target int) int {
	sum := 0
	for _, n := range arr {
		sum += n
	}
	if sum < target || ((target&1)^(sum&1)) != 0 {
		return 0
	}
	return subset2(arr, (target+sum)>>1)
}

func subset2(nums []int, s int) int {
	if s < 0 {
		return 0
	}
	dp := make([]int, s+1)
	dp[0] = 1
	for _, n := range nums {
		for i := s; i >= n; i-- {
			dp[i] += dp[i-n]
		}
	}
	return dp[s]
}
