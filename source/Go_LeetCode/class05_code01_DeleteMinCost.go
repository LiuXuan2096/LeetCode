package Go_LeetCode

import (
	"sort"
	"strings"
)

/*
题目描述：
给定两个字符串s1和s2，问s2最少删除多少字符可以成为s1的子串？
比如 s1 = "abcde"，s2 = "axbc"，s2删掉'x'即可，返回1
*/

// 解法一
// 求出str2所有的子序列，然后按照长度排序，长度大的排在前面。
// 然后考察哪个子序列字符串和s1的某个子串相等(KMP)，答案就出来了。
// 分析：
// 因为题目原本的样本数据中，有特别说明s2的长度很小。所以这么做也没有太大问题，也几乎不会超时。
// 但是如果某一次考试给定的s2长度远大于s1，这么做就不合适了。
func MinCost_05_01(s1, s2 string) int {
	s2Sub := make([]string, 0)
	process_05_01([]rune(s2), 0, "", &s2Sub)
	sort.Slice(s2Sub, func(i, j int) bool {
		return len(s2Sub[i]) >= len(s2Sub[j])
	})
	for _, str := range s2Sub {
		if strings.Contains(s1, str) {
			return len(s2) - len(str)
		}
	}
	return len(s2)
}

func process_05_01(str2 []rune, index int, path string, list *[]string) {
	if index == len(str2) {
		*list = append(*list, path)
		return
	}
	process_05_01(str2, index+1, path, list)
	process_05_01(str2, index+1, path+string(str2[index]), list)
}

// 解法二：动态规划版本的解法
func MinCost_dp(s1, s2 string) int {
	c1 := []rune(s1)
	c2 := []rune(s2)
	dp := make([][]int, len(c2)+1)
	for i := range dp {
		dp[i] = make([]int, len(c1)+1)
	}
	// dp[i][j] 的含义 :
	// s2中前缀i长度的字符串，至少删除多少个字符可以变成
	//s1中前缀j长度字符串的 后缀串
	// s2前缀0长度，删掉0个字符，可以变成s1前缀任意长度字符串的 后缀串
	// 所以dp[0][....] = 0，所以省略了
	for i := 1; i <= len(c2); i++ {
		// s2前缀i长度，需要都删掉，可以变成s1前缀0长度字符串的 后缀串
		dp[i][0] = i
		for j := 1; j <= len(c1); j++ {
			// 如果 c2[i-1] == c1[j-1]
			// 可能性1 要么 c2[i-1] 和 c1[j-1] 进行匹配
			// 问题回到 dp[i-1][j-1]
			// 可能性2 要么 c2[i-1] 和 c1[j-1] 不进行匹配
			// s2[0...i] 依然删除 s2[i] 问题回到 dp[i-1][j] + 1
			// dp[i][j] = Math.min(dp[i-1][j-1], dp[i-1][j] + 1);
			// 实际上
			// 只需要考虑可能性1
			// 例
			// abc 去匹配 abc.....c
			// 删掉 .....c 和 删掉 c..... 是一个意思
			if c2[i-1] == c1[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 如果 c2[i-1] != c1[j-1]
				// c2[i-1] 和 c1[j-1] 不进行匹配
				// s2[0...i] 删除 s2[i] 问题回到 dp[i-1][j] + 1
				dp[i][j] = dp[i-1][j] + 1
			}
		}
	}
	ans := dp[len(c2)][0]
	for j := 1; j <= len(c1); j++ {
		if dp[len(c2)][j] < ans {
			ans = dp[len(c2)][j]
		}
	}
	return ans
}
