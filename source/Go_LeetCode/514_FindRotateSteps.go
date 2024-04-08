package Go_LeetCode

import "math"

func findRotateSteps(ring string, key string) int {
	r := []rune(ring)
	N := len(r)
	mapPositions := make(map[rune][]int)
	for i, ch := range r {
		mapPositions[ch] = append(mapPositions[ch], i)
	}
	str := []rune(key)
	M := len(str)
	dp := make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, M+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return process_514(0, 0, str, mapPositions, N, dp)
}

// 电话里：指针指着的上一个按键preButton
// 目标里：此时要搞定哪个字符？keyIndex
// map : key 一种字符 value: 哪些位置拥有这个字符
// N: 电话大小
// f(0, 0, aim, map, N)
func process_514(preButton, keyIndex int, str []rune, mapPositions map[rune][]int, N int, dp [][]int) int {
	if dp[preButton][keyIndex] != -1 {
		return dp[preButton][keyIndex]
	}
	ans := math.MaxInt32
	if keyIndex == len(str) {
		ans = 0
	} else {
		cur := str[keyIndex]
		// 还有字符需要搞定呢！
		nextPositions := mapPositions[cur]
		for _, next := range nextPositions {
			cost := dial(preButton, next, N) + 1 + process_514(next, keyIndex+1, str, mapPositions, N, dp)
			ans = min_514(ans, cost)
		}
	}
	dp[preButton][keyIndex] = ans
	return ans
}

func dial(i1, i2, size int) int {
	return min_514(abs_514(i1-i2), min_514(i1, i2)+size-max_514(i1, i2))
}

func min_514(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max_514(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func abs_514(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
