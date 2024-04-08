package Go_LeetCode

import "sort"

// 给定一个数组arr，代表每个人的能力值。再给定一个非负数k。
// 如果两个人能力差值正好为k，那么可以凑在一起比赛，
//一局比赛只有两个人
// 返回最多可以同时有多少场比赛
func MaxPairNum2(arr []int, k int) int {
	if k < 0 || arr == nil || len(arr) < 2 {
		return 0
	}
	sort.Ints(arr)
	ans := 0
	N := len(arr)
	L, R := 0, 0
	usedR := make([]bool, N)
	for L < N && R < N {
		if usedR[L] {
			L++
		} else if L >= R {
			R++
		} else {
			distance := arr[R] - arr[L]
			if distance == k {
				ans++
				usedR[R] = true
				R++
				L++
			} else if distance < k {
				R++
			} else {
				L++
			}
		}
	}
	return ans
}
