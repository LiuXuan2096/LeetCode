package Go_LeetCode

import "sort"

// 可改为动态规划的暴力递归解法
//func twoCitySchedCost(costs [][]int) int {
//	if costs == nil || len(costs) < 2 || len(costs)%2 != 0 {
//		return 0
//	}
//	N := len(costs)
//	M := N >> 1
//	return process_1029(costs, 0, M)
//}
//
//// i位置到最后位置所有的人，往A和B区域分配！
//// A区域还有rest个名额!
//// 返回把i位置到最后的人分配万，分配完，
////并且最终A和B区域同样多的情况下，i...最后这些人，
////整体成本最低是多少！
//func process_1029(costs [][]int, i, rest int) int {
//	if i == len(costs) {
//		return 0
//	}
//	if len(costs)-i == rest {
//		return costs[i][0] + process_1029(costs, i+1, rest-1)
//	}
//	if rest == 0 {
//		return costs[i][1] + process_1029(costs, i+1, rest)
//	}
//	// 当前面试者，可以去A，或者去B
//	p1 := costs[i][0] + process_1029(costs, i+1, rest-1)
//	p2 := costs[i][1] + process_1029(costs, i+1, rest)
//	if p1 <= p2 {
//		return p1
//	}
//	return p2
//
//}

// 贪心解法
func twoCitySchedCost(costs [][]int) int {
	N := len(costs)
	arr := make([]int, N)
	sum := 0
	for i := 0; i < N; i++ {
		arr[i] = costs[i][1] - costs[i][0]
		sum += costs[i][0]
	}
	sort.Ints(arr)
	M := N >> 1
	for i := 0; i < M; i++ {
		sum += arr[i]
	}
	return sum
}
