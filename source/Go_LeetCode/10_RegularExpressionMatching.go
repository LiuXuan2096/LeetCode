package Go_LeetCode

/*
 * 这个题的暴力递归版本的解法，从这个解法改动态规划

func isMatch(s string, p string) bool {
	if s == "" || p == "" {
		return false
	}

	str := []rune(s)
	pattern := []rune(p)

	return isValid(str, pattern) && process1(str, pattern, 0, 0)
}

func process1(str, pattern []rune, si, pi int) bool {
	// 判断str从si位置开始的字符串 即str[i:] 是否能被从pi位置开始的模式串匹配 pattern[pi:]
	// 隐含的条件pattern[pi] != '*' 因为已经通过了 isValid 函数的检验
	// 情况一：si越界了
	if si == len(str) {
		if pi == len(pattern) {
			return true // pi也越界了，说明匹配完成
		}
		if pi+1 < len(pattern) && pattern[pi+1] == '*' {
			// 说明pattern还没有匹配结束，且pattern的下一位是 ‘*’
			// 则pattern[pi]和pattern[pi+1]可以匹配出一个空串 ”“
			// 那么继续判断后续的pattern是否能匹配出一个空串
			return process1(str, pattern, si, pi+2)
		}
		return false // 到这里说明匹配失败
	}

	// 情况二：si没越界时，pi越界了，即pattern已经匹配结束了
	if pi == len(pattern) {
		return si == len(str)
	}

	// 情况三：si pi都没越界，且 pattern[pi+1] != '*'
	if pi+1 >= len(pattern) || pattern[pi+1] != '*' {
		return (str[si] == pattern[pi] || pattern[pi] == '.') && process1(str, pattern, si+1, pi+1)
	}

	// 情况四：si pi都没越界 且pi+1是‘*’（能执行到这说明pi+1位置是‘*’）
	if pattern[pi] != '.' && str[si] != pattern[pi] {
		return process1(str, pattern, si, pi+2)
	}

	// 执行到这里说明 si pi都没越界，且pi+1是‘*’ 且pi位置可匹配si位置
	if process1(str, pattern, si, pi+2) {
		return true
	}

	// 能执行到这潜台词是 pi+1=‘*’
	for si < len(str) && (str[si] == pattern[pi] || pattern[pi] == '.') {
		if process1(str, pattern, si+1, pi+2) { // 所以这里是pi+2
			return true
		}
		si++
	}
	return false
}

func isValid(str []rune, pattern []rune) bool {
	for _, ch := range str {
		if ch == '.' || ch == '*' {
			return false
		}
	}
	for i := 0; i < len(pattern); i++ {
		if pattern[i] == '*' && i == 0  {
			return false
		}
	}
	return true
}
*/
func isMatch(s string, p string) bool {

}
