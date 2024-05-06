package Go_LeetCode

import "math"

func cherryPickup(grid [][]int) int {
	N := len(grid)
	M := len(grid[0])
	dp := make([][][]int, N)
	for i := range dp {
		dp[i] = make([][]int, M)
		for j := range dp[i] {
			dp[i][j] = make([]int, N)
			for k := range dp[i][j] {
				dp[i][j][k] = math.MinInt32
			}
		}
	}
	ans := process_741(grid, 0, 0, 0, dp)
	if ans < 0 {
		return 0
	}
	return ans
}

// process_741 这个递归过程是模拟两个人同时从起点走
// A(x1, y1)  B (x2, x1+y1-x2)
// A B两人同时走则根据题意，它们坐标之和一定相同
// 一个人表示去时的走法 一个人表示回来的走法 两种走法加起来能摘到的最大殷桃数量
func process_741(grid [][]int, x1, y1, x2 int, dp [][][]int) int {
	N, M := len(grid), len(grid[0])
	if x1 == N || y1 == M || x2 == N || x1+y1-x2 == M {
		return math.MinInt32
	}
	if dp[x1][y1][x2] != math.MinInt32 {
		return dp[x1][y1][x2]
	}
	if x1 == N-1 && y1 == M-1 {
		dp[x1][y1][x2] = grid[x1][y1]
		return dp[x1][y1][x2]
	}
	next := math.MinInt32
	// A B 都向下走
	next = max(next, process_741(grid, x1+1, y1, x2+1, dp))
	// A向下走 B向右走
	next = max(next, process_741(grid, x1+1, y1, x2, dp))
	// A B 都向右走
	next = max(next, process_741(grid, x1, y1+1, x2, dp))
	// A向右走 B向下走
	next = max(next, process_741(grid, x1, y1+1, x2+1, dp))

	if grid[x1][y1] == -1 || grid[x2][x1+y1-x2] == -1 || next == -1 {
		dp[x1][y1][x2] = -1
		return -1
	}
	if x1 == x2 {
		dp[x1][y1][x2] = grid[x1][y1] + next
	} else {
		dp[x1][y1][x2] = grid[x1][y1] + grid[x2][x1+y1-x2] + next
	}
	return dp[x1][y1][x2]
}
