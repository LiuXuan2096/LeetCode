package Go_LeetCode

import "math"

// 所有的货，重量和价值，都在w和v数组里
// 为了方便，其中没有负数
// bag背包容量，不能超过这个载重
// 返回：不超重的情况下，能够得到的最大价值
//
//func MaxValue(w []int, v []int, bag int) int {
//	if w == nil || v == nil || len(w) != len(v) || len(w) == 0 {
//		return 0
//	}
//	// 尝试函数
//	return process(w, v, 0, bag)
//}
//
//func process(w []int, v []int, index, rest int) int {
//	// index 0~N
//	// rest 负~bag
//	if rest < 0 {
//		return -1 // 说明上一步的选择无效
//	}
//	if index == len(w) {
//		return 0
//	}
//	var p1 = process(w, v, index+1, rest) // 没有选当前物品
//	var p2 = 0
//	var next = process(w, v, index+1, rest-w[index])
//	if next != -1 {
//		p2 = next + v[index]
//	}
//	return int(math.Max(float64(p1), float64(p2)))
//}

func KnapSackDP(w []int, v []int, bag int) int {
	if w == nil || v == nil || len(w) != len(v) || len(w) == 0 {
		return 0
	}
	var N = len(w)
	var dp = make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, bag+1)
	}
	for i := N - 1; i >= 0; i-- {
		for rest := 0; rest <= bag; rest++ {
			var p1 = dp[i+1][rest]
			var p2 = 0
			var next int
			if rest-w[i] < 0 {
				next = -1
			} else {
				next = dp[i+1][rest-w[i]]
			}
			if next != -1 {
				p2 = v[i] + next
			}
			dp[i][rest] = int(math.Max(float64(p1), float64(p2)))
		}
	}
	return dp[0][bag]
}
