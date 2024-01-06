package Go_LeetCode

import "strconv"

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	if subRoot == nil {
		return true
	}
	big := preSerial(root)
	small := preSerial(subRoot)

	return getIndexof2(big, small) != -1
}

func preSerial(root *TreeNode) []string {
	var ans []string
	pres(root, &ans)
	return ans
}

func pres(root *TreeNode, ans *[]string) {
	if root == nil {
		(*ans) = append(*ans, "#")
	} else {
		(*ans) = append(*ans, strconv.Itoa(root.Val))
		pres(root.Left, ans)
		pres(root.Right, ans)
	}
}

func getIndexof2(match, pattern []string) int {
	var matchIndex = 0
	var patternIndex = 0
	var next = getNextArray3(pattern)
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

func getNextArray3(pattern []string) []int {
	if len(pattern) == 1 {
		return []int{-1}
	}
	var next = make([]int, len(pattern))
	next[0] = -1
	next[1] = 0
	var curIndex = 2
	var curPrefixEnd = 0
	for curIndex < len(pattern) {
		if pattern[curIndex-1] == pattern[curPrefixEnd] {
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
