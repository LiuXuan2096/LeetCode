package Go_LeetCode

// 递归版本的做法

func sameTypeSameNumber(str1, str2 []rune) bool {
	if len(str1) != len(str2) {
		return false
	}
	mapTable := make(map[rune]int)
	for _, char := range str1 {
		mapTable[char]++
	}
	for _, char := range str2 {
		mapTable[char]--
		if mapTable[char] < 0 {
			return false
		}
	}
	return true
}

//func isScramble(s1 string, s2 string) bool {
//	if (s1 == "" && s2 != "") || (s1 != "" && s2 == "") {
//		return false
//	}
//	if s1 == "" && s2 == "" {
//		return true
//	}
//	if s1 == s2 {
//		return true
//	}
//	str1 := []rune(s1)
//	str2 := []rune(s2)
//	if !sameTypeSameNumber(str1, str2) {
//		return false
//	}
//	N := len(s1)
//	return process_87(str1, str2, 0, 0, N)
//}
//
//// 返回str1[从L1开始往右长度为size的子串]和str2[从L2开始往右长度为size的子串]是否互为旋变字符串
//// 在str1中的这一段和str2中的这一段一定是等长的，所以只用一个参数size
//func process_87(str1, str2 []rune, L1, L2, size int) bool {
//	if size == 1 {
//		return str1[L1] == str2[L2]
//	} else {
//		// 枚举每一种情况，有一个计算出互为旋变就返回true。都算不出来最后返回false
//		for leftPart := 1; leftPart < size; leftPart++ {
//			// 可能性1：如果1左对2左，并且1右对2右
//			p1 := process_87(str1, str2, L1, L2, leftPart) && process_87(str1, str2, L1+leftPart, L2+leftPart, size-leftPart)
//			// 可能性2:如果1左对2右，并且1右对2左
//			p2 := process_87(str1, str2, L1, L2+size-leftPart, leftPart) && process_87(str1, str2, L1+leftPart, L2, size-leftPart)
//			if p1 || p2 {
//				return true
//			}
//		}
//	}
//	return false
//}

// 记忆化搜索的版本，时间复杂度已经和动态规划版本的一样了
func isScramble(s1 string, s2 string) bool {
	if (s1 == "" && s2 != "") || (s1 != "" && s2 == "") {
		return false
	}
	if s1 == "" && s2 == "" {
		return true
	}
	if s1 == s2 {
		return true
	}
	str1 := []rune(s1)
	str2 := []rune(s2)
	if !sameTypeSameNumber(str1, str2) {
		return false
	}
	N := len(s1)
	dp := make([][][]int, N)
	// dp[i][j][k] = 0 processDP(i,j,k)状态之前没有算过的
	// dp[i][j][k] = -1 processDP(i,j,k)状态之前算过的,返回值是false
	// dp[i][j][k] = 1 processDP(i,j,k)状态之前算过的,返回值是true
	for i := range dp {
		dp[i] = make([][]int, N)
		for j := range dp[i] {
			dp[i][j] = make([]int, N+1)
		}
	}
	return process_87_dp(str1, str2, 0, 0, N, dp)
}

func process_87_dp(str1, str2 []rune, L1, L2, size int, dp [][][]int) bool {
	if dp[L1][L2][size] != 0 {
		return dp[L1][L2][size] == 1
	}
	var ans bool
	if size == 1 {
		ans = str1[L1] == str2[L2]
	} else {
		for leftPart := 1; leftPart < size; leftPart++ {
			p1 := process_87_dp(str1, str2, L1, L2, leftPart, dp) && process_87_dp(str1, str2, L1+leftPart, L2+leftPart, size-leftPart, dp)
			p2 := process_87_dp(str1, str2, L1, L2+size-leftPart, leftPart, dp) && process_87_dp(str1, str2, L1+leftPart, L2, size-leftPart, dp)
			if p1 || p2 {
				ans = true
				break
			}
		}
	}
	if ans {
		dp[L1][L2][size] = 1
	} else {
		dp[L1][L2][size] = -1
	}
	return ans
}
