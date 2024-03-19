package from_forceRecursive_to_DP

// 题目链接：// 测试链接：https://leetcode.com/problems/longest-palindromic-subsequence/

// 暴力递归版本的解法
func lplll(s string) int {
	if s == "" {
		return 0
	}
	str := []rune(s)
	return f(str, 0, len(str)-1)
}

func f(str []rune, L, R int) int {
	if L == R {
		return 1
	}
	if L == R-1 {
		if str[L] == str[R] {
			return 2
		}
		return 1
	}
	p1 := f(str, L+1, R-1)
	p2 := f(str, L+1, R)
	p3 := f(str, L, R-1)
	p4 := 0
	if str[L] == str[R] {
		p4 = 2 + f(str, L+1, R-1)
	}
	return max(max(p1, p2), max(p3, p4))
}

// 动态规划版本的解法
func longestPalindromeSubseq(s string) int {
	if s == "" {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	str := []rune(s)
	N := len(str)
	dp := make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, N)
	}

	dp[N-1][N-1] = 1
	for i := 0; i < N-1; i++ {
		dp[i][i] = 1
		if str[i] == str[i+1] {
			dp[i][i+1] = 2
		} else {
			dp[i][i+1] = 1
		}
	}

	for i := N - 3; i >= 0; i-- {
		for j := i + 2; j < N; j++ {
			dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			if str[i] == str[j] {
				dp[i][j] = max(dp[i][j], dp[i+1][j-1]+2)
			}
		}
	}
	return dp[0][N-1]
}
