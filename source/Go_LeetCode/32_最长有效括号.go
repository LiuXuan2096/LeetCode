package Go_LeetCode

func longestValidParentheses(s string) int {
	if s == "" || len(s) < 2 {
		return 0
	}
	str := []rune(s)
	// dp[i] : 子串必须以i位置结尾的情况下，往左最远能扩出多长的有效区域
	dp := make([]int, len(str))
	var pre int
	var ans int
	for i := 1; i < len(str); i++ {
		// 当前谁和i位置的)，去配！
		// 与str[i]配对的左括号的位置 pre
		if str[i] == ')' {
			pre = i - dp[i-1] - 1
			if pre >= 0 && str[pre] == '(' {
				if pre == 0 {
					dp[i] = dp[i-1] + 2
				} else {
					dp[i] = dp[i-1] + 2 + dp[pre-1]
				}
			}
		}
		ans = max_32(ans, dp[i])
	}
	return ans
}

func max_32(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
