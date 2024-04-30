package Go_LeetCode

func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}

	ans := 0 //在0时刻卖出时的最大利润就是0 所以ans初始为0
	min := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > ans {
			ans = prices[i] - min
		}
	}
	return ans
}
