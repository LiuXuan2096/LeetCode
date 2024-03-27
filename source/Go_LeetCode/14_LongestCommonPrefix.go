package Go_LeetCode

import "math"

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}
	chs := []rune(strs[0])
	minVal := math.MaxInt32
	for _, str := range strs {
		tmp := []rune(str)
		index := 0
		for index < len(tmp) && index < len(chs) {
			if chs[index] != tmp[index] {
				break
			}
			index++
		}
		minVal = min(index, minVal)
		if minVal == 0 {
			return ""
		}
	}
	return string(chs[:minVal])
}

//func min(a, b int) int {
//	if a <= b {
//		return a
//	}
//	return b
//}
