package from_forceRecursive_to_DP

/*
题目描述：
背包问题
给定两个长度都为N的数组weights和values，weights[i]和values[i]分别代表 i号物品的重量和价值
给定一个正数bag，表示一个载重bag的袋子，装的物品不能超过这个重量
返回能装下的最大价值
*/

// MaxValue 最原始的暴力递归版本的解法
// 所有的货，重量和价值，都在w和v数组里
// 为了方便，其中没有负数
// bag背包容量，不能超过这个载重
// 返回：不超重的情况下，能够得到的最大价值
func MaxValue(w, v []int, bag int) int {
	if w == nil || v == nil || len(w) != len(v) || len(w) == 0 {
		return 0
	}
	return process(w, v, 0, bag)
}

// index:0-N
// rest <= bag
func process(w, v []int, index, rest int) int {
	if rest < 0 {
		return -1 // 表示当前选择使背包剩余容量为负，返回-1说明当前选择是个无效解
	}
	if index == len(w) {
		return 0
	}
	p1 := process(w, v, index+1, rest)
	p2 := 0
	next := process(w, v, index+1, rest-w[index])
	if next != -1 {
		p2 = v[index] + next
	}
	return max(p1, p2)
}

// MaxValueDP 动态规划版本的解法
func MaxValueDP(w []int, v []int, bag int) int {
	if w == nil || v == nil || len(w) != len(v) || len(w) == 0 {
		return 0
	}
	N := len(w)
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, bag+1)
	}

	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= bag; rest++ {
			p1 := dp[index+1][rest]
			p2 := 0
			if rest-w[index] >= 0 {
				next := dp[index+1][rest-w[index]]
				p2 = v[index] + next
			}
			dp[index][rest] = max(p1, p2)
		}
	}
	return dp[0][bag]
}
