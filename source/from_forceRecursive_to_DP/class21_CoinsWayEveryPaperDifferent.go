package from_forceRecursive_to_DP

/*
题目描述：arr是货币数组，其中的值都是正数。再给定一个正数aim。
每个值都认为是一张货币，
即便是值相同的货币也认为每一张都是不同的，
返回组成aim的方法数
例如：arr = {1,1,1}，aim = 2
第0个和第1个能组成2，第1个和第2个能组成2，第0个和第2个能组成2
一共就3种方法，所以返回3
*/

func CoinWays(arr []int, aim int) int {
	return process_CoinsWays(arr, 0, aim)
}

// arr[index....] 组成正好rest这么多的钱，有几种方法
func process_CoinsWays(arr []int, index, rest int) int {
	if rest < 0 {
		return 0
	}
	if index == len(arr) {
		if rest == 0 {
			return 1
		}
		return 0
	}
	return process_CoinsWays(arr, index+1, rest) + process_CoinsWays(arr, index+1, rest-arr[index])
}

func CoinWays_dp(arr []int, aim int) int {
	if aim == 0 {
		return 1
	}
	N := len(arr)
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}
	dp[N][0] = 1

	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= aim; rest++ {
			dp[index][rest] = dp[index+1][rest]
			if rest-arr[index] >= 0 {
				dp[index][rest] += dp[index+1][rest-arr[index]]
			}
		}
	}
	return dp[0][aim]
}
