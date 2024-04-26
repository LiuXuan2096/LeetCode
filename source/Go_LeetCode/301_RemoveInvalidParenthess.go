package Go_LeetCode

func removeInvalidParentheses(s string) []string {
	ans := []string{}
	remove_301(s, &ans, 0, 0, []byte{'(', ')'})
	return ans
}

// modifyIndex <= checkIndex
// 只查s[checkIndex....]的部分，因为之前的一定已经调整对了
// 但是之前的部分是怎么调整对的，调整到了哪？就是modifyIndex
// 比如：
// ( ( ) ( ) ) ) ...
// 0 1 2 3 4 5 6
// 一开始当然checkIndex = 0，modifyIndex = 0
// 当查到6的时候，发现不对了，
// 然后可以去掉2位置、4位置的 )，都可以
// 如果去掉2位置的 ), 那么下一步就是
// ( ( ( ) ) ) ...
// 0 1 2 3 4 5 6
// checkIndex = 6 ，modifyIndex = 2
// 如果去掉4位置的 ), 那么下一步就是
// ( ( ) ( ) ) ...
// 0 1 2 3 4 5 6
// checkIndex = 6 ，modifyIndex = 4
// 也就是说，
// checkIndex和modifyIndex，分别表示查的开始 和 调的开始，之前的都不用管了  par  (  )
func remove_301(s string, ans *[]string, checkIndex, deleteIndex int, par []byte) {
	count := 0
	for i := checkIndex; i < len(s); i++ {
		if s[i] == par[0] {
			count++
		}
		if s[i] == par[1] {
			count--
		}
		// i check计数<0的第一个位置
		if count < 0 {
			for j := deleteIndex; j <= i; j++ {
				if s[j] == par[1] && (j == deleteIndex || s[j-1] != par[1]) {
					newStr := s[:j] + s[j+1:]
					remove_301(newStr, ans, i, j, par)
				}
			}
			return
		}
	}
	reversed := reverseString_301(s)
	if par[0] == '(' {
		remove_301(reversed, ans, 0, 0, []byte{')', '('})
	} else {
		*ans = append(*ans, reversed)
	}
}

func reverseString_301(s string) string {
	str := []rune(s)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	return string(str)
}
