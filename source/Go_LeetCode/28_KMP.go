package Go_LeetCode

func SstrStr(haystack string, needle string) int {
	if len(haystack) == 0 || len(needle) == 0 || len(haystack) < len(needle) {
		return -1
	}
	var match = []byte(haystack)
	var pattern = []byte(needle)
	var matchIndex = 0
	var patternIndex = 0
	// O(M) m<=n
	var next = getNextArray(pattern)
	// O(N)
	for matchIndex < len(match) && patternIndex < len(pattern) {
		if match[matchIndex] == pattern[patternIndex] {
			matchIndex++
			patternIndex++
		} else if next[patternIndex] == -1 {
			matchIndex++
		} else {
			patternIndex = next[patternIndex]
		}
	}
	if patternIndex == len(pattern) {
		return matchIndex - patternIndex
	} else {
		return -1
	}
}

func getNextArray(pattern []byte) []int {
	if len(pattern) == 1 {
		return []int{-1}
	}
	var next = make([]int, len(pattern))
	next[0] = -1
	next[1] = 0
	var i = 2        // 目前在哪个位置上求next数组的值
	var curIndex = 0 // 当前是哪个位置的值再和i-1位置的字符比较
	for i < len(pattern) {
		if pattern[curIndex] == pattern[i-1] {
			next[i] = curIndex + 1
			i++
			curIndex++
		} else if curIndex > 0 {
			curIndex = next[curIndex]
		} else {
			next[i] = 0
			i++
		}
	}
	return next
}
