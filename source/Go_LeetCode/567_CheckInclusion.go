package Go_LeetCode

func checkInclusion(s1 string, s2 string) bool {
	tmp := s2
	s2 = s1
	s1 = tmp

	if s1 == "" || s2 == "" || len(s1) < len(s2) {
		return false
	}
	str2 := []rune(s2)
	M := len(str2)
	count := make([]int, 256)
	for i := 0; i < M; i++ {
		count[str2[i]]++
	}
	all := M
	str1 := []rune(s1)
	R := 0
	// 0~M-1
	// 最早的M个字符，让其窗口初步形成
	for ; R < M; R++ {
		if count[str1[R]] > 0 {
			all--
		}
		count[str1[R]]--
	}
	// 窗口初步形成了，并没有判断有效无效，决定下一个位置一上来判断
	// 接下来的过程，窗口右进一个，左吐一个
	for ; R < len(str1); R++ {
		if all == 0 {
			return true
		}
		if count[str1[R]] > 0 {
			all--
		}
		count[str1[R]]--
		if count[str1[R-M]] >= 0 {
			// 这个位置即将离开窗口，如果这个位置的字符在s2里有
			// 在这个字符离开窗口后 应该给all+1
			all++
		}
		count[str1[R-M]]++
	}
	if all == 0 {
		return true
	}
	return false
}
