package from_forceRecursive_to_DP

import "math"

/*
arr是面值数组，其中的值都是正数且没有重复。再给定一个正数aim。
每个值都认为是一种面值，且认为张数是无限的。
返回组成aim的最少货币数
*/
func MinCoinsNoLimit(arr []int, aim int) int {
	return process_minCoinsNoLimit(arr, 0, aim)
}

// arr[index...]面值，每种面值张数自由选择，
// 搞出rest正好这么多钱，返回最小张数
// 拿math.MaxInt32标记怎么都搞定不了
func process_minCoinsNoLimit(arr []int, index, rest int) int {
	if index == len(arr) {
		if rest == 0 {
			return 0
		}
		return math.MaxInt32
	}

	ans := math.MaxInt32
	for zhang := 0; zhang*arr[index] <= rest; zhang++ {
		next := process_minCoinsNoLimit(arr, index+1, rest-zhang*arr[index])
		if next != math.MaxInt32 {
			ans = min(ans, zhang+next)
		}
	}
	return ans
}

func MinCoinNoLimit_dp(arr []int, aim int) int {
	if aim == 0 {
		return 0
	}
	N := len(arr)
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}
	dp[N][0] = 0

	for j := 1; j <= aim; j++ {
		dp[N][j] = math.MaxInt32
	}

	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {
			dp[index][rest] = dp[index+1][rest]
			if rest-arr[index] >= 0 && dp[index][rest-arr[index]] != math.MaxInt32 {
				dp[index][rest] = min(dp[index][rest], dp[index][rest-arr[index]]+1)
			}
		}
	}
	return dp[0][aim]
}
