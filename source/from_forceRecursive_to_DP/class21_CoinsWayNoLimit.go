package from_forceRecursive_to_DP

func CoinsWay(arr []int, aim int) int {
	if arr == nil || len(arr) == 0 || aim < 0 {
		return 0
	}
	return process_CoinsWaysNoLimit(arr, 0, aim)
}

// arr[index....] 所有的面值，每一个面值都可以任意选择张数，
//组成正好rest这么多钱，方法数多少？
func process_CoinsWaysNoLimit(arr []int, index, rest int) int {
	if index == len(arr) {
		if rest == 0 {
			return 1
		}
		return 0
	}
	ways := 0
	for zhang := 0; zhang*arr[index] <= rest; zhang++ {
		ways += process_CoinsWaysNoLimit(arr, index+1, rest-(arr[index]*zhang))
	}
	return ways
}

func CoinsWayNolimit_dp(arr []int, aim int) int {
	if arr == nil || len(arr) == 0 || aim < 0 {
		return 0
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
				dp[index][rest] += dp[index][rest-arr[index]]
			}
		}

	}
	return dp[0][aim]
}
