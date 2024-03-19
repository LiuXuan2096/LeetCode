package Go_LeetCode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	records := [26]int{}
	for index := 0; index < len(s); index++ {
		if s[index] == t[index] {
			continue
		}
		sCharIndex := s[index] - 'a'
		records[sCharIndex]++
		tCharIndex := t[index] - 'a'
		records[tCharIndex]--
	}
	for _, record := range records {
		if record != 0 {
			return false
		}
	}
	return true
}
