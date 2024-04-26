package Go_LeetCode

func kInversePairs(n int, k int) int {
	if n < 1 || k < 0 {
		return 0
	}
	mod := 1000000007
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	dp[0][0] = 1

	for i := 1; i <= n; i++ {
		dp[i][0] = 1
		for j := 1; j <= k; j++ {
			dp[i][j] = (dp[i][j-1] + dp[i-1][j]) % mod
			if j >= i {
				dp[i][j] = (dp[i][j] - dp[i-1][j-i] + mod) % mod
			}
		}
	}
	return dp[n][k]
}
