package Go_LeetCode

// 暴力递归版本
//func canIWin(maxChoosableInteger int, desiredTotal int) bool {
//	if desiredTotal == 0 {
//		return true
//	}
//	if (maxChoosableInteger * (maxChoosableInteger + 1) / 2) < desiredTotal {
//		// 因为不能重复拿，如果可选择的数组累加和小于desiredTotal 那么永远不可能累加到desiredTotal
//		return false
//	}
//	arr := make([]int, maxChoosableInteger)
//	for i := 0; i < maxChoosableInteger; i++ {
//		arr[i] = i + 1
//	}
//	return process_464(arr, desiredTotal)
//}
//
//// 当前轮到先手拿，
//// 先手只能选择在arr中还存在的数字，
//// 还剩rest这么值，
//// 返回先手会不会赢
//func process_464(arr []int, rest int) bool {
//	if rest <= 0 {
//		return false
//	}
//	for i := 0; i < len(arr); i++ {
//		if arr[i] != -1 {
//			cur := arr[i]
//			arr[i] = -1
//			next := process_464(arr, rest-cur)
//			arr[i] = cur
//			if !next {
//				return true
//			}
//		}
//	}
//	return false
//}

// 从暴力递归版本进化为压缩状态的动态规划
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if desiredTotal == 0 {
		return true
	}
	if maxChoosableInteger*(maxChoosableInteger+1)/2 < desiredTotal {
		return false
	}
	dp := make([]int, 1<<(maxChoosableInteger+1))
	// dp[status] == 1  true
	// dp[status] == -1  false
	// dp[status] == 0  process(status) 没算过！去算！
	return process_464(maxChoosableInteger, 0, desiredTotal, dp)
}

// 为什么明明status和rest是两个可变参数，却只用status来代表状态(也就是dp)
// 因为选了一批数字之后，得到的和一定是一样的，
//所以rest是由status决定的，所以rest不需要参与记忆化搜索
func process_464(choose, status, rest int, dp []int) bool {
	if dp[status] != 0 {
		return dp[status] == 1
	}
	var ans bool
	if rest > 0 {
		for i := 1; i <= choose; i++ {
			if (1<<i)&status == 0 {
				if !process_464(choose, (status | (1 << i)), rest-i, dp) {
					ans = true
					break
				}
			}
		}
	}
	if ans {
		dp[status] = 1
	} else {
		dp[status] = -1
	}
	return ans
}
