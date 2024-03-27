package Go_LeetCode

func strStr(haystack string, needle string) int {
	if haystack == "" || needle == "" || len(needle) < 1 || len(haystack) < len(needle) {
		return -1
	}
	str1 := []rune(haystack)
	str2 := []rune(needle)
	x := 0
	y := 0
	next := getNextArray_28(str2)
	for x < len(str1) && y < len(str2) {
		if str1[x] == str2[y] {
			x++
			y++
		} else if next[y] == -1 {
			x++
		} else {
			y = next[y]
		}
	}
	if y == len(str2) {
		return x - y
	}
	return -1
}

func getNextArray_28(str2 []rune) []int {
	if len(str2) == 1 {
		return []int{-1}
	}
	next := make([]int, len(str2))
	next[0] = -1
	next[1] = 0
	i := 2  // 目前在哪个位置上求next数组的值
	cn := 0 // 当前是哪个位置的值再和i-1位置的字符比较
	for i < len(next) {
		if str2[i-1] == str2[cn] {
			cn++
			next[i] = cn
			i++
		} else if cn > 0 {
			cn = next[cn]
		} else {
			next[i] = 0
			i++
		}
	}
	return next
}
