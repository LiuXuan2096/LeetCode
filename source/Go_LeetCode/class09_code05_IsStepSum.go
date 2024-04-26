package Go_LeetCode

/*
题目描述：
定义何为step sum？比如680，680 + 68 + 6 = 754，680的step sum叫754。
给定一个正数num，判断它是不是某个数的step sum
*/

func IsStepNum(stepSum int) bool {
	L := 0
	R := stepSum
	var M, cur int
	for L <= R {
		M = L + ((R - L) >> 1)
		cur = stepSumFunc(M)
		if cur == stepSum {
			return true
		} else if cur < stepSum {
			L = M + 1
		} else {
			R = M - 1
		}
	}
	return false
}

func stepSumFunc(num int) int {
	sum := 0
	for num != 0 {
		sum += num
		num /= 10
	}
	return sum
}
