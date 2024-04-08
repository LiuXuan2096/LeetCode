package Go_LeetCode

/*
 * 只由小写字母（a~z）组成的一批字符串，都放在字符类型的数组String[] arr中，
 * 如果其中某两个字符串，所含有的字符种类完全一样，就将两个字符串算作一类 比如：baacba和bac就算作一类
 * 虽然长度不一样，但是所含字符的种类完全一样（a、b、c） 返回arr中有多少类？
 *
 */

func Type2(arr []string) int {
	types := make(map[int]bool)
	for _, str := range arr {
		key := 0
		for _, ch := range str {
			key |= 1 << (ch - 'a')
		}
		types[key] = true
	}
	return len(types)
}
