package Go_LeetCode

func minDistance(word1 string, word2 string) int {
	return minCost_72(word1, word2, 1, 1, 1)
}

// minCost_72 ic dc rc分别表示插入 删除 和替换这三种操作的代价
func minCost_72(s1 string, s2 string, ic, dc, rc int) int {
	if s1 == s2 {
		return 0
	}
	str1 := []rune(s1)
	str2 := []rune(s2)
	N := len(str1) + 1
	M := len(str2) + 1

	// 样本对应模型
	// 存储动态规划的过程的表，dp[i][j]表示s1[...i]位置转换成
	// s2[...j]位置的代价
	dp := make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, M)
	}
	for i := 1; i < N; i++ {
		dp[i][0] = dc * i
	}
	for j := 1; j < M; j++ {
		dp[0][j] = ic * j
	}
	/*
				dp[i][j] = dp[i - 1][j - 1] + (str1[i - 1] == str2[j - 1] ? 0 : rc);
			dp[i][j] = Math.min(dp[i][j], dp[i][j - 1] + ic);
			dp[i][j] = Math.min(dp[i][j], dp[i - 1][j] + dc);
		}
	*/
	for i := 1; i < N; i++ {
		for j := 1; j < M; j++ {
			if str1[i-1] == str2[j-1] {
				// 情况1 s1 i位置的字符 和s2 j位置的字符相同，则：
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 否则，情况2：
				dp[i][j] = dp[i-1][j-1] + rc
			}
			// 情况三：s1 0-i位置的字符转换为 s2 0-j-1位置的字符的代价 加上 s2插入1个字符的代价
			dp[i][j] = min_72(dp[i][j], dp[i][j-1]+ic)
			// 情况四：s1 0-i-1位置的字符转换为s2 0-j位置的字符的代价 加上 s1删除一个字符的代价
			dp[i][j] = min_72(dp[i][j], dp[i-1][j]+dc)
		}
	}
	return dp[N-1][M-1]
}

func min_72(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
