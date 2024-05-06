package Go_LeetCode

func distinctSubseqII(s string) int {
	if s == "" {
		return 0
	}
	const m = 1000000007
	str := []rune(s)
	count := make([]int, 26) //存放当前最新的以每个字符结尾时的子序列字面值的数量
	all := 1                 // all表示当前有的子序列字面值的数量 空集算一个
	for _, x := range str {
		// 加入一个新的字符后新增的子序列数量是all 即把原有的结果每个结尾加上一个新的字符x
		// 但是其中重复计算的字面值数量为count[x-'a']
		add := (all - count[x-'a'] + m) % m
		all = (all + add) % m
		count[x-'a'] = (count[x-'a'] + add) % m
	}
	all = (all - 1 + m) % m
	return all

}
