package from_forceRecursive_to_DP

/*
题目描述：
规定1和A对应、2和B对应、3和C对应...26和Z对应
那么一个数字字符串比如"111”就可以转化为:
"AAA"、"KA"和"AK"
给定一个只有数字字符组成的字符串str，返回有多少种转化结果
*/

// 最原始的暴力递归版本的解法
func Number(str string) int {
	if str == "" {
		return 0
	}
	return process_C19_2([]rune(str), 0)
}

// process_C19_2
// str[0..i-1]转化无需过问
// str[i.....]去转化，返回有多少种转化方法
func process_C19_2(str []rune, i int) int {
	if i == len(str) {
		return 1
	}
	// i 没到最后 说明还有剩余字符没转化
	if str[i] == '0' {
		// 说明之前的决策有问题，‘0'字符不应该出现在第一个
		return 0
	}
	// 两种决策：i位置单转 和 i位置与i+1位置一起转
	ways := process_C19_2(str, i+1)
	if i+1 < len(str) && (str[i]-'0')*10+(str[i+1]-'0') < 27 {
		ways += process_C19_2(str, i+2)
	}
	return ways
}

// NumberDP 动态规划版本的解法
func NumberDP(s string) int {
	if s == "" {
		return 0
	}
	str := []rune(s)
	N := len(str)
	dp := make([]int, N+1)
	dp[N] = 1
	for i := N - 1; i >= 0; i-- {
		if str[i] != '0' {
			ways := dp[i+1]
			if i+1 < N && (str[i]-'0')*10+(str[i+1]-'0') < 27 {
				ways += dp[i+2]
			}
			dp[i] = ways
		}
	}
	return dp[0]
}
