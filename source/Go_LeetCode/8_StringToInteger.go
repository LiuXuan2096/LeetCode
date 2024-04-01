package Go_LeetCode

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	if s == "" {
		return 0
	}

	s = removeHeadZero(strings.TrimSpace(s))

	if s == "" {
		return 0
	}
	str := []rune(s)

	if !isValid_8(str) {
		return 0
	}

	posi := str[0] != '-'      // 标志这个数是否是负数
	minq := math.MinInt32 / 10 // 系统最小除以10
	minr := math.MinInt32 % 10 // 系统最小模10
	res := 0
	cur := 0

	for i := func() int {
		if str[0] == '-' || str[0] == '+' {
			return 1
		}
		return 0
	}(); i < len(str); i++ {
		cur = int('0' - str[i])
		if res < minq || (res == minq && cur < minr) {
			// 执行到这里 说明转换后的数字溢出了
			if posi {
				return math.MaxInt32
			}
			return math.MinInt32
		}
		res = res*10 + cur
	}

	if posi && res == math.MinInt32 {
		return math.MaxInt32
	}

	if posi {
		return -res
	}
	return res
}

func removeHeadZero(str string) string {
	// 判断字符串开头是否有 ‘+’ 或者 ‘-’ 如果有的话
	// 从index位置为1 开始去除’0‘
	r := strings.HasPrefix(str, "+") || strings.HasPrefix(str, "-")
	s := func() int {
		if r {
			return 1
		}
		return 0
	}()
	start := s
	for ; s < len(str); s++ {
		if str[s] != '0' {
			break
		}
	} // 此时 s是左边第一个非‘0’字符的索引

	// 从后向前遍历，去掉非数字字符
	e := -1
	for i := len(str) - 1; i >= start; i-- {
		if str[i] < '0' || str[i] > '9' {
			e = i
		}
	} // e 到了最左边的（即下标最小的）非数字字符的索引

	var addOrMinus = ""
	if start == 1 {
		addOrMinus = string(str[0])
	}
	if e == -1 {
		return addOrMinus + str[s:]
	} else {
		return addOrMinus + str[s:e]
	}
}

func isValid_8(chas []rune) bool {
	if chas[0] != '-' && chas[0] != '+' && (chas[0] < '0' || chas[0] > '9') {
		return false
	}
	if (chas[0] == '-' || chas[0] == '+') && len(chas) == 1 {
		return false
	}

	for i := 1; i < len(chas); i++ {
		if chas[i] < '0' || chas[i] > '9' {
			return false
		}
	}

	return true
}
