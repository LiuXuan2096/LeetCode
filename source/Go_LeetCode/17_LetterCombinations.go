package Go_LeetCode

var phone = [][]rune{
	{'a', 'b', 'c'},
	{'d', 'e', 'f'},
	{'g', 'h', 'i'},
	{'j', 'k', 'l'},
	{'m', 'n', 'o'},
	{'p', 'q', 'r', 's'},
	{'t', 'u', 'v'},
	{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	ans := []string{}
	if len(digits) == 0 {
		return ans
	}
	str := []rune(digits)
	path := make([]rune, len(str))
	process_17(str, 0, path, &ans)
	return ans
}

func process_17(str []rune, i int, path []rune, ans *[]string) {
	if i == len(str) {
		*ans = append(*ans, string(path))
	} else {
		cands := phone[str[i]-'2']
		for _, cur := range cands {
			path[i] = cur
			process_17(str, i+1, path, ans)
		}
	}
}
