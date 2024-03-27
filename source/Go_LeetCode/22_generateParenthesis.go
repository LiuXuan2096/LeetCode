package Go_LeetCode

func generateParenthesis(n int) []string {
	path := make([]rune, n*2)
	var ans []string
	process_22(path, 0, 0, n, &ans)
	return ans
}

func process_22(path []rune, i, leftMinusRight, leftRest int, ans *[]string) {
	if i == len(path) {
		*ans = append(*ans, string(path))
	} else {
		if leftRest > 0 {
			path[i] = '('
			process_22(path, i+1, leftMinusRight+1, leftRest-1, ans)
		}
		if leftMinusRight > 0 {
			path[i] = ')'
			process_22(path, i+1, leftMinusRight-1, leftRest, ans)
		}
	}
}
