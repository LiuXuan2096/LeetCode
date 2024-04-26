package Go_LeetCode

// 这个题没什么很统一的像比如 动态规划那样的套路
// 这个题属于 知道解法就很容易做 不知道解法就很难做出来的那种题
// 算是用贪心做的 知道做法即可 不要去想证明
func findMinMoves(machines []int) int {
	if machines == nil || len(machines) == 0 {
		return 0
	}
	size := len(machines)
	sum := 0
	for i := 0; i < size; i++ {
		sum += machines[i]
	}
	if sum%size != 0 {
		return -1
	}
	avg := sum / size
	leftSum := 0
	ans := 0
	for i := 0; i < len(machines); i++ {
		// 下面的左侧、右侧都是相对于i而言的 i的左侧和右侧
		leftRest := leftSum - i*avg                                 //左侧当前有的衣服 - 左侧需要的衣服 就是当前左侧还缺多少衣服
		rightRest := (sum - leftSum - machines[i]) - (size-i-1)*avg //右侧当前有的衣服 - 右侧需要的衣服 就是当前右侧还缺多少衣服
		if leftRest < 0 && rightRest < 0 {
			ans = max_517(ans, abs_517(leftRest)+abs_517(rightRest))
		} else {
			ans = max_517(ans, max_517(abs_517(leftRest), abs_517(rightRest)))
		}
		leftSum += machines[i]
	}
	return ans
}

func max_517(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func abs_517(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
