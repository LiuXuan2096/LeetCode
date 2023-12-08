package Go_LeetCode

import "math"

func LongestPalindromeSubseq(s string) int {
	if len(s) == 0 {
		return 0
	}
	str := []byte(s)
	return process5(str, 0, len(s)-1)
}

func process5(str []byte, L, R int) int {
	if L == R {
		return 1
	}
	if L == R-1 {
		if str[L] == str[R] {
			return 2
		} else {
			return 1
		}
	}
	var p1 = process5(str, L+1, R-1)
	var p2 = process5(str, L, R-1)
	var p3 = process5(str, L+1, R)
	var p4 = 0
	if str[L] == str[R] {
		p4 = 2 + process5(str, L+1, R-1)
	}
	return int(math.Max(math.Max(float64(p1), float64(p2)), math.Max(float64(p3), float64(p4))))
}

func longestPalindromeSubseq(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	str := []byte(s)
	var N = len(s)
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
			dp[i][j] = int(math.Max(float64(dp[i][j-1]), float64(dp[i+1][j])))
			if str[i] == str[j] {
				dp[i][j] = int(math.Max(float64(dp[i][j]), float64(dp[i+1][j-1]+2)))
			}
		}
	}
	return dp[0][N-1]
}
