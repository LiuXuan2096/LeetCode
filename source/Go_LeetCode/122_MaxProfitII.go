package Go_LeetCode

func maxProfitII(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}

	ans := 0
	for i := 1; i < len(prices); i++ {
		if diff := prices[i] - prices[i-1]; diff > 0 {
			ans += diff
		}
	}

	return ans
}
