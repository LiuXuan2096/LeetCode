package Go_LeetCode

func repeatedSubstringPattern(s string) bool {
	length := len(s)
	if length == 1 {
		return false
	}
	next := getNextArray_459(s)
	pattern := s
	match := s + s
	x := 0
	y := 1
	for x < len(pattern) && y < len(match)-1 {
		if pattern[x] == match[y] {
			x++
			y++
		} else if next[x] != -1 {
			x = next[x]
		} else {
			y++
		}
	}
	if x == len(pattern) {
		return true
	} else {
		return false
	}
}

func getNextArray_459(str string) []int {
	chars := []byte(str)
	next := make([]int, len(chars))
	next[0] = -1
	next[1] = 0
	i := 2
	cn := 0
	for i < len(next) {
		if chars[i-1] == chars[cn] {
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
