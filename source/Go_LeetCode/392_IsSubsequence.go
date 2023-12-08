package Go_LeetCode

func IsSubsequence(s string, t string) bool {
	if len(s) == 0 && len(t) != 0 {
		return true
	}
	if len(t) == 0 && len(s) != 0 {
		return false
	}
	if len(t) == 0 && len(s) == 0 {
		return true
	}

	var source []byte = []byte(t)
	var template []byte = []byte(s)
	var tLen = len(template)

	var i = 0
	var j = 0

	for ; i < len(source); i++ {
		if source[i] == template[j] {
			j++
			if j == tLen {
				return true
			}
		}
	}
	return false
}
