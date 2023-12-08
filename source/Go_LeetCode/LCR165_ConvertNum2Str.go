package Go_LeetCode

import "strconv"

// CrackNumber 从右往左的动态规划
// 就是递归方法的动态规划版本
// dp[i]表示：str[i...]有多少种转化方式
func CrackNumber(ciphertext int) int {
	var str = strconv.Itoa(ciphertext)
	var numArray []byte = []byte(str)
	var dp []int = make([]int, len(numArray)+1)
	var N = len(numArray)
	dp[N] = 1
	for i := N - 1; i >= 0; i-- {
		var ways = dp[i+1]
		if str[i] != '0' {

			if i+1 < len(str) && ((str[i]-'0')*10+str[i+1]-'0') < 26 {
				ways += dp[i+2]
			}
			dp[i] = ways
		} else {
			dp[i] = ways
		}
	}
	return dp[0]
}

//func process(str []byte, i int) int {
//	if i == len(str) {
//		return 1
//	}
//	// i没到最后 说明还有字符
//	if str[i] == '0' {
//		return 0
//	}
//	// 选择一 只转换i位置的数字
//	var ways = process(str, i+1)
//	// 选择二 将i和i+1位置的数字转换为一个字母
//	if i+1 < len(str) && ((str[i]-'0')*10+str[i+1]-'0') < 26 {
//		ways += process(str, i+2)
//	}
//	return ways
//}
