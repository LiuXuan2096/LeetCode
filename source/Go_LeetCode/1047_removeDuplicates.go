package Go_LeetCode

func removeDuplicates(s string) string {
	var result string
	for _, ch := range s {
		if len(result) == 0 || result[len(result)-1] != byte(ch) {
			result += string(ch)
		} else {
			result = result[:len(result)-1]
		}
	}
	return result
}
