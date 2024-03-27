package Go_LeetCode

func reverseStr(s string, k int) string {
	ss := []byte(s)
	length := len(s)
	for i := 0; i < length; i += 2 * k {
		// 1. 每隔2k个字符的前k个字符进行反转
		// 2. 剩余字符小于2k 但大于等于k个 则反转前k个字符
		if i+k <= length {
			reverse_541(ss[i : i+k])
		} else {
			reverse_541(ss[i:length])
		}
	}
	return string(ss)
}

func reverse_541(str []byte) {
	left := 0
	right := len(str) - 1
	for left < right {
		str[left], str[right] = str[right], str[left]
		left++
		right--
	}
}
