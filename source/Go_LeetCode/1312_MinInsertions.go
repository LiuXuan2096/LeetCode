package Go_LeetCode

import "math"

func minInsertions(s string) int {
	if s == "" || len(s) < 2 {
		return 0
	}
	str := []rune(s)
	N := len(str)

	// 范围尝试模型
	// 动态规划表的的含义：
	// dp[i][j] 表示s[i...j]范围的字符 添加成回文串的最少操作次数
	dp := make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, N)
	}
	// 为动态规划表的[i][i+1]位置赋值
	for i := 0; i < N-1; i++ {
		if str[i] == str[i+1] {
			dp[i][i+1] = 0
		} else {
			dp[i][i+1] = 1
		}
	}
	for i := N - 3; i >= 0; i-- {
		for j := i + 2; j < N; j++ {
			// 可能性1：s[i...j-1]已经成为回文串，j位置再加一个字符
			// 那么s[i...j]位置都成了回文串
			p1 := dp[i][j-1] + 1
			// 可能性2：s[i+1...j]已经成为回文串,i位置再加一个字符
			p2 := dp[i+1][j] + 1
			// 可能性3: s[i]=s[j]
			p3 := math.MaxInt32
			if str[i] == str[j] {
				p3 = dp[i+1][j-1]
			}
			dp[i][j] = min_1312(p1, min_1312(p2, p3))
		}
	}
	return dp[0][N-1]
}

func min_1312(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
