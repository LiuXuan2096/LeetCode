package Go_LeetCode

import "math"

type Info_378 struct {
	near int
	num  int
}

// noMoreNum 矩阵中<=value的数有几个
func noMoreNum(matrix [][]int, value int) Info_378 {
	near := math.MinInt32
	num := 0
	N := len(matrix)
	M := len(matrix[0])
	row := 0
	col := M - 1

	for row < N && col >= 0 {
		if matrix[row][col] <= value {
			if matrix[row][col] > near {
				near = matrix[row][col]
			}
			num += col + 1
			row++
		} else {
			col--
		}
	}
	return Info_378{near, num}

}

// kthSmallest 二分查找
func kthSmallest(matrix [][]int, k int) int {
	N := len(matrix)
	if N == 0 {
		return 0
	}
	M := len(matrix[0])
	left := matrix[0][0]      //最小的
	right := matrix[N-1][M-1] //最大的
	ans := 0

	for left <= right {
		mid := left + ((right - left) >> 1)
		info := noMoreNum(matrix, mid)
		if info.num < k {
			left = mid + 1
		} else {
			ans = info.near
			right = mid - 1
		}
	}
	return ans
}
