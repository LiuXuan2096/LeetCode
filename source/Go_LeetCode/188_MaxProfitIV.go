package Go_LeetCode

func maxProfitIV(k int, prices []int) int {
	if prices == nil || len(prices) == 0 || k < 1 {
		return 0
	}

	N := len(prices)
	if k >= N/2 {
		return allTrans(prices)
	}

	dp := make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	// 初始化dp[...][0] = 0
	for i := 0; i < N; i++ {
		dp[i][0] = 0
	}
	// 处理dp[1][j]
	for j := 1; j <= k; j++ {
		// 先填好dp[1][j]的值
		p1 := dp[0][j]
		p2 := dp[1][j-1] + prices[1] - prices[1]
		p3 := dp[0][j-1] + prices[1] - prices[0]
		dp[1][j] = max(p1, max(p2, p3))
		best := max(dp[1][j-1]-prices[1], dp[0][j-1]-prices[0]) //用来简化后续的枚举
		// dp[1][j] 准备好一些枚举，接下来准备好的枚举
		for i := 2; i < N; i++ {
			cur1 := dp[i-1][j]
			newP := dp[i][j-1] - prices[i]
			best = max(newP, best)
			dp[i][j] = max(cur1, best+prices[i])
		}
	}

	return dp[N-1][k]
}

func allTrans(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}
