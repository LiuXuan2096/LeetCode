package Go_LeetCode

func reverseWords(s string) string {
	//1.使用双指针删除冗余的空格
	slowIndex, fastIndex := 0, 0
	str := []byte(s)
	// 删除头部冗余空格
	for len(str) > 0 && fastIndex < len(str) && str[fastIndex] == ' ' {
		fastIndex++
	}
	// 删除单词间冗余空格
	for ; fastIndex < len(str); fastIndex++ {
		if fastIndex-1 > 0 && str[fastIndex] == str[fastIndex-1] && str[fastIndex] == ' ' {
			continue
		}
		str[slowIndex] = str[fastIndex]
		slowIndex++
	}
	// 删除尾部冗余空格
	if slowIndex-1 > 0 && str[slowIndex-1] == ' ' {
		str = str[:slowIndex-1]
	} else {
		str = str[:slowIndex]
	}
	//2.反转整个字符串
	reverse_151(str)
	//3.反转单个单词， i 单词开始位置  j 单词结束位置
	i := 0
	for i < len(str) {
		j := i
		for ; j < len(str) && str[j] != ' '; j++ {

		}
		reverse_151(str[i:j])
		i = j
		i++
	}
	return string(str)
}

func reverse_151(chars []byte) {
	left := 0
	right := len(chars) - 1
	for left < right {
		chars[left], chars[right] = chars[right], chars[left]
		left++
		right--
	}
}
