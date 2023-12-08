package Go_LeetCode

func Change(amount int, coins []int) int {
	return process6(coins, 0, amount)
}

// arr[index....] 所有的面值，每一个面值都可以任意选择张数，组成正好rest这么多钱，方法数多少？
func process6(coins []int, index int, rest int) int {
	if index == len(coins) {
		if rest == 0 {
			return 1
		} else {
			return 0
		}
	}
	var ways = 0
	for zhang := 0; zhang*coins[index] <= rest; zhang++ {
		ways += process6(coins, index+1, rest-(zhang*coins[index]))
	}
	return ways
}

func change(amount int, coins []int) int {
	var N = len(coins)
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}
	//  dp[i][j]表示从下标为i...N-1的数组中选货币 组成值为j的方案有多少种
	dp[N][0] = 1
	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= amount; rest++ {
			dp[index][rest] = dp[index+1][rest]
			if rest-coins[index] >= 0 {
				dp[index][rest] += dp[index][rest-coins[index]]
			}
		}
	}
	return dp[0][amount]
}
