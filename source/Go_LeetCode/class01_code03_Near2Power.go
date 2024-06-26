package Go_LeetCode

/*
给定一个非负整数num，如何不用循环语句，
返回>=num，并且离num最近的，2的某次方
*/
func Near2Power(n int) int {
	n--
	n = n | (n >> 1)
	n = n | (n >> 2)
	n = n | (n >> 4)
	n = n | (n >> 8)
	n = n | (n >> 16)
	if n < 0 {
		return 1
	}
	return n + 1
}
