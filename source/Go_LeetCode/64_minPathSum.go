package Go_LeetCode

import "math"

func minPathSum(grid [][]int) int {
	var row = len(grid)
	var col = len(grid[0])
	dp := make([]int, col)
	dp[0] = grid[0][0]
	for j := 1; j < col; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}
	for i := 1; i < row; i++ {
		dp[0] += grid[i][0]
		for j := 1; j < col; j++ {
			dp[j] = int(math.Min(float64(dp[j-1]), float64(dp[j]))) + grid[i][j]
		}
	}
	return dp[col-1]
}
