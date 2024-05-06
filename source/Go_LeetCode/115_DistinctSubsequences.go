package Go_LeetCode

func NumDistinct(s string, t string) int {

	chars_s := []rune(s)
	chars_t := []rune(t)
	N := len(chars_s)
	M := len(chars_t)
	// S[...i]的所有子序列中，包含多少个字面值等于T[...j]这个字符串的子序列
	// 记为dp[i][j]
	// 可能性1）S[...i]的所有子序列中，都不以s[i]结尾，则dp[i][j]肯定包含dp[i-1][j]
	// 可能性2）S[...i]的所有子序列中，都必须以s[i]结尾，
	// 这要求S[i] == T[j]，则dp[i][j]包含dp[i-1][j-1]
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, M)
	}
	if chars_s[0] == chars_t[0] {
		dp[0][0] = 1
	} else {
		dp[0][0] = 0
	}

	for i := 1; i < N; i++ {
		if s[i] == t[0] {
			dp[i][0] = dp[i-1][0] + 1
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}

	for i := 1; i < N; i++ {
		for j := 1; j <= min(i, M-1); j++ {
			dp[i][j] = dp[i-1][j]
			if s[i] == t[j] {
				dp[i][j] += dp[i-1][j-1]
			}
		}
	}

	return dp[N-1][M-1]
}
