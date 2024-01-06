package Go_LeetCode

func RrotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	var goalgoal = goal + goal
	return getIndexOf(goalgoal, s) != -1
}

func getIndexOf(match, pattern string) int {
	if len(match) < len(pattern) {
		return -1
	}
	var matchArr = []byte(match)
	var patternArr = []byte(pattern)
	var matchIndex = 0
	var patternIndex = 0
	next := getNextArray2(patternArr)
	for matchIndex < len(matchArr) && patternIndex < len(patternArr) {
		if matchArr[matchIndex] == patternArr[patternIndex] {
			patternIndex++
			matchIndex++
		} else if next[patternIndex] == -1 {
			matchIndex++
		} else {
			patternIndex = next[patternIndex]
		}
	}
	if patternIndex == len(patternArr) {
		return matchIndex - patternIndex
	} else {
		return -1
	}
}

func getNextArray2(patternArr []byte) []int {
	if len(patternArr) == 1 {
		return []int{-1}
	}
	var next = make([]int, len(patternArr))
	next[0] = -1
	next[1] = 0
	var curIndex = 2
	var curPrefixEnd = 0
	for curIndex < len(patternArr) {
		if patternArr[curIndex-1] == patternArr[curPrefixEnd] {
			next[curIndex] = curPrefixEnd + 1
			curIndex++
			curPrefixEnd++
		} else if curPrefixEnd > 0 {
			curPrefixEnd = next[curPrefixEnd]
		} else {
			next[curIndex] = 0
			curIndex++
		}
	}
	return next
}
