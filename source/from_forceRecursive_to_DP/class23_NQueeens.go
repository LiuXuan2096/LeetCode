package from_forceRecursive_to_DP

/*
N皇后问题是指在N*N的棋盘上要摆N个皇后，
要求任何两个皇后不同行、不同列， 也不在同一条斜线上
给定一个整数n，返回n皇后的摆法有多少种。n=1，返回1
n=2或3，2皇后和3皇后问题无论怎么摆都不行，返回0
n=8，返回92
*/

func NQueens(n int) int {
	if n < 1 {
		return 0
	}
	record := make([]int, n)
	return process_NQueens(0, record, n)
}

func process_NQueens(i int, record []int, n int) int {
	if i == n {
		return 1
	}
	res := 0
	for j := 0; j < n; j++ {
		if isValid(record, i, j) {
			record[i] = j
			res += process_NQueens(i+1, record, n)
		}
	}
	return res
}

func isValid(record []int, i, j int) bool {
	for k := 0; k < i; k++ {
		if j == record[k] || abs(record[k]-j) == abs(k-i) {
			return false
		}

	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
