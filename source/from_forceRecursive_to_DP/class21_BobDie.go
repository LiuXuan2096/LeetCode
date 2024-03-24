package from_forceRecursive_to_DP

import "math"

/*
给定5个参数，N，M，row，col，k
表示在N*M的区域上，醉汉Bob初始在(row,col)位置
Bob一共要迈出k步，且每步都会等概率向上下左右四个方向走一个单位
任何时候Bob只要离开N*M的区域，就直接死亡
返回k步之后，Bob还在N*M的区域的概率
*/

func LivePosibility(row, col, k, N, M int) float64 {
	return float64(process_BobDie(row, col, k, N, M)) / math.Pow(4, float64(k))
}

// 目前在row，col位置，还有rest步要走，
//走完了如果还在棋盘中就获得1个生存点，返回总的生存点数
func process_BobDie(row, col, rest, N, M int) int {
	if row < 0 || row == N || col < 0 || col == M {
		return 0
	}
	// 还在棋盘中！
	if rest == 0 {
		return 1
	}
	// 还在棋盘中！还有步数要走
	up := process_BobDie(row-1, col, rest-1, N, M)
	down := process_BobDie(row+1, col, rest-1, N, M)
	left := process_BobDie(row, col-1, rest-1, N, M)
	right := process_BobDie(row, col+1, rest-1, N, M)
	return up + down + left + right
}
