package Go_LeetCode

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	str := []rune(s)
	var mapArray [256]int //存放当前遇到的字符上一次出现的位置
	for i := 0; i < 256; i++ {
		mapArray[i] = -1
	}
	// 收集答案
	result := 0
	pre := -1 // i-1位置结尾的情况下，往左推，推不动的位置的索引
	curLen := 0
	for i := 0; i < len(str); i++ {
		// i位置结尾的情况下，往左推，推不动的位置的索引
		// 更新pre的值
		pre = max(pre, mapArray[str[i]])
		curLen = i - pre
		result = max(curLen, result)
		mapArray[str[i]] = i
	}
	return result
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
