package Go_LeetCode

/*
这个题目前的代码版本是错误的 无法通过全部测试用例
*/

//func FindPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
//	dp := make([][][]int, m)
//	for i := range dp {
//		dp[i] = make([][]int, n)
//		for j := range dp[i] {
//			dp[i][j] = make([]int, maxMove+1)
//		}
//	}
//	for i := 0; i < m; i++ {
//		for j := 0; j < n; j++ {
//			dp[i][j][0] = 1
//		}
//	}
//	for rest := 1; rest <= maxMove; rest++ {
//		for r := 0; r < m; r++ {
//			for c := 0; c < n; c++ {
//				dp[r][c][rest] = pick(dp, m, n, r-1, c, rest-1)
//				dp[r][c][rest] = pick(dp, m, n, r+1, c, rest-1)
//				dp[r][c][rest] = pick(dp, m, n, r, c-1, rest-1)
//				dp[r][c][rest] = pick(dp, m, n, r, c-1, rest-1)
//			}
//		}
//	}
//	return int(math.Pow(4, float64(maxMove))) - dp[startRow][startColumn][maxMove]
//}
//
//func pick(dp [][][]int, N, M, r, c, rest int) int {
//	if r < 0 || r == N || c < 0 || c == M {
//		return 0
//	}
//	return dp[r][c][rest]
//}

//func FindPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
//	return process7(startRow, startColumn, maxMove, m, n)
//}
//
//func process7(row, col, rest, N, M int) int {
//	if row < 0 || row == N || col < 0 || col == M {
//		// 走到棋盘外了
//		return 1
//	}
//	if rest == 0 {
//		// 步数走完了 还在棋盘里
//		return 0
//	}
//	// 还在棋盘中 且 步数没走完
//	var up = process7(row-1, col, rest-1, N, M)
//	var down = process7(row+1, col, rest-1, N, M)
//	var left = process7(row, col-1, rest-1, N, M)
//	var right = process7(row, col+1, rest-1, N, M)
//	return up + down + left + right
//}

func FindPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, maxMove+1)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j][0] = 0
		}
	}
	for rest := 1; rest <= maxMove; rest++ {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				dp[i][j][rest] = pick(dp, m, n, i-1, j, rest-1)
				dp[i][j][rest] += pick(dp, m, n, i+1, j, rest-1)
				dp[i][j][rest] += pick(dp, m, n, i, j-1, rest-1)
				dp[i][j][rest] += pick(dp, m, n, i, j+1, rest-1)
			}
		}
	}
	return dp[startRow][startColumn][maxMove] % (10e9 + 7)
}

func pick(dp [][][]int, m, n, r, c, rest int) int {
	if r < 0 || r == m || c < 0 || c == n {
		// 走到棋盘外了
		return 1
	}
	return dp[r][c][rest]
}
