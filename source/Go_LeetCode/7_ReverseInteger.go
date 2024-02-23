package Go_LeetCode

import "math"

func reverse(x int) int {
	neg := ((x >> 31) & 1) == 1 // 通过位运算判断neg是否是负数
	x = func() int {
		if neg {
			return x
		}
		return -x
	}()

	m := math.MinInt32 / 10 // 系统最小除以10
	o := math.MinInt32 % 10 // 系统最小模10
	res := 0

	for x != 0 {
		if res < m || (res == m && x%10 < o) {
			// 执行到这里说明res的范围会溢出，按照题意返回0
			return 0
		}
		res = res*10 + x%10
		x /= 10
	}

	if neg {
		return res
	}
	return int(math.Abs(float64(res)))
}
