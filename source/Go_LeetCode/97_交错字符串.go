package Go_LeetCode

func isInterleave(s1 string, s2 string, s3 string) bool {
	if s1 == "" && s2 == "" && s3 == "" {
		return true
	} else if s1 == "" || s2 == "" || s3 == "" {
		if s3 == "" {
			return false
		} else if s1 == "" {
			return s2 == s3
		} else {
			return s1 == s3
		}
	}
	str1 := []rune(s1)
	str2 := []rune(s2)
	str3 := []rune(s3)
	if len(str3) != len(str1)+len(str2) {
		return false
	}
	dp := make([][]bool, len(str1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(str2)+1)
	}
	dp[0][0] = true
	for i := 1; i <= len(str1); i++ {
		if str1[i-1] != str3[i-1] {
			break
		}
		dp[i][0] = true
	}
	for j := 1; j <= len(str2); j++ {
		if str2[j-1] != str3[j-1] {
			break
		}
		dp[0][j] = true
	}
	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if (str1[i-1] == str3[i+j-1] && dp[i-1][j]) || (str2[j-1] == str3[i+j-1] && dp[i][j-1]) {
				dp[i][j] = true
			}
		}
	}
	return dp[len(str1)][len(str2)]
}
