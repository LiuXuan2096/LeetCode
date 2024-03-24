package from_forceRecursive_to_DP

/*
给定一个正数数组arr，
请把arr中所有的数分成两个集合，尽量让两个集合的累加和接近
返回最接近的情况下，较小集合的累加和
*/
func SplitSumClosed(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return process_SplitSum(arr, 0, sum/2)
}

// arr[i...]可以自由选择，请返回累加和尽量接近rest，
// 但不能超过rest的情况下，最接近的累加和是多少？
func process_SplitSum(arr []int, index, rest int) int {
	if index == len(arr) {
		return 0
	} else {
		// 还有数 arr【i】
		// 可能性1：不使用arr[i]
		p1 := process_SplitSum(arr, index+1, rest)
		// 可能性2: 使用arr[i]
		p2 := 0
		if arr[index] <= rest {
			p2 = arr[index] + process_SplitSum(arr, index+1, rest-arr[index])
		}
		return max(p1, p2)
	}
}

func SplitSum_dp(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	sum := 0
	for _, num := range arr {
		sum += num
	}
	sum /= 2
	N := len(arr)
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, sum+1)
	}
	for i := N - 1; i >= 0; i-- {
		for rest := 0; rest <= sum; rest++ {
			// 可能性1 不使用arr[i]
			p1 := dp[i+1][rest]
			// 可能性2 使用arr[i]
			p2 := 0
			if arr[i] <= rest {
				p2 = arr[i] + dp[i+1][rest-arr[i]]
			}
			dp[i][rest] = max(p1, p2)
		}
	}
	return dp[0][sum]
}
