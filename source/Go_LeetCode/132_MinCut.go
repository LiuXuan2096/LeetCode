package Go_LeetCode

import "math"

func minCut(s string) int {
	if s == "" || len(s) < 2 {
		return 0
	}
	str := []rune(s)
	N := len(str)
	checkMap := createCheckMap(str, N)

	// 动态规划表中 dp[i]的含义是 字符串s在范围[i...N-1]
	// 范围上 有几个回文串
	dp := make([]int, N+1)
	dp[N] = 0
	for i := N - 1; i >= 0; i-- {
		if checkMap[i][N-1] {
			dp[i] = 1
		} else {
			next := math.MaxInt32 - 1
			for j := i; j < N; j++ {
				if checkMap[i][j] {
					next = min_132(next, dp[j+1])
				}
			}
			dp[i] = next + 1
		}
	}
	return dp[0] - 1
}

func createCheckMap(str []rune, N int) [][]bool {
	ans := make([][]bool, N)
	for i := range ans {
		ans[i] = make([]bool, N)
	}
	for i := 0; i < N-1; i++ {
		ans[i][i] = true
		ans[i][i+1] = str[i] == str[i+1]
	}
	ans[N-1][N-1] = true
	for i := N - 3; i >= 0; i-- {
		for j := i + 2; j < N; j++ {
			ans[i][j] = str[i] == str[j] && ans[i+1][j-1]
		}
	}
	return ans
}

func min_132(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
