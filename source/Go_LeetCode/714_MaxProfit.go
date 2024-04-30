package Go_LeetCode

func maxProfitVI(prices []int, fee int) int {
	if prices == nil || len(prices) < 2 {
		return 0
	}

	N := len(prices)
	// 初始化最佳买入状态（考虑手续费）
	bestBuy := -prices[0] - fee
	// 初始化最佳卖出状态
	bestSell := 0

	for i := 1; i < N; i++ {
		// 当前买入（考虑手续费）
		// 来到i位置了！
		// 如果在i必须买  收入 - 批发价 - fee
		curBuy := max(bestSell-prices[i]-fee, bestBuy)
		// 如果在i必须卖  整体最优（收入 - 良好批发价 - fee）
		curSell := bestBuy + prices[i]
		// 更新最佳买入和卖出状态
		bestBuy = curBuy
		bestSell = max(bestSell, curSell)
	}

	return bestSell
}
