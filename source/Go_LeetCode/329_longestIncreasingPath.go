package Go_LeetCode

//func longestIncreasingPath(matrix [][]int) int {
//	ans := 0
//	N := len(matrix)
//	M := len(matrix[0])
//	for i := 0; i < N; i++ {
//		for j := 0; j < M; j++ {
//			ans = max(ans, process_329(matrix, i, j))
//		}
//	}
//	return ans
//}
//
//func process_329(matrix [][]int, i, j int) int {
//	up := 0
//	if i > 0 && matrix[i][j] < matrix[i-1][j] {
//		up = process_329(matrix, i-1, j)
//	}
//	down := 0
//	if i < len(matrix)-1 && matrix[i][j] < matrix[i+1][j] {
//		down = process_329(matrix, i+1, j)
//	}
//	left := 0
//	if j > 0 && matrix[i][j] < matrix[i][j-1] {
//		left = process_329(matrix, i, j-1)
//	}
//	right := 0
//	if j < len(matrix[0])-1 && matrix[i][j] < matrix[i][j+1] {
//		right = process_329(matrix, i, j+1)
//	}
//	return max(max(up, down), max(left, right)) + 1
//}

func longestIncreasingPath(matrix [][]int) int {
	ans := 0
	N := len(matrix)
	M := len(matrix[0])
	dp := make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, M)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			ans = max(ans, process_329(matrix, i, j, dp))
		}
	}
	return ans
}

func process_329(matrix [][]int, i, j int, dp [][]int) int {
	if dp[i][j] != 0 {
		return dp[i][j]
	}
	up := 0
	if i > 0 && matrix[i][j] < matrix[i-1][j] {
		up = process_329(matrix, i-1, j, dp)
	}
	down := 0
	if i < len(matrix)-1 && matrix[i][j] < matrix[i+1][j] {
		down = process_329(matrix, i+1, j, dp)
	}
	left := 0
	if j > 0 && matrix[i][j] < matrix[i][j-1] {
		left = process_329(matrix, i, j-1, dp)
	}
	right := 0
	if j < len(matrix[0])-1 && matrix[i][j] < matrix[i][j+1] {
		right = process_329(matrix, i, j+1, dp)
	}
	ans := max(max(up, down), max(left, right)) + 1
	dp[i][j] = ans
	return ans
}
