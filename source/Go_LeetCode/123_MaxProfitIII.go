package Go_LeetCode

/*
ans：用于存储当前计算出的最大利润。
doneOnceMinusBuyMax：表示在某一天卖出股票后（已经做过一次交易），手中剩余的现金减去最初购买股票花费的金额后的最大值。初始值为 -prices[0]，即假设第一天买入股票。
doneOnceMax：表示在某一天卖出股票后（已经做过一次交易），手中剩余的现金的最大值。初始值为0。
min：用于存储到目前为止看到的最低股票价格。初始值设置为 prices[0]，即第一天的股票价格。
*/

func maxProfitIII(prices []int) int {
	if prices == nil || len(prices) < 2 {
		return 0
	}
	N := len(prices)
	k := 2
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
