package Go_LeetCode

import "math"

func countDigitOne(n int) int {
	if n < 1 {
		return 0
	}
	// num -> 13625
	// len = 5位数
	length := getLenOfNum(n)
	if length == 1 {
		return 1
	}
	// num 13625
	// tmp1 10000
	//
	// num 7872328738273
	// tmp1 1000000000000
	tmp1 := powerBaseOf10(length - 1)
	// num最高位 num / tmp1
	first := n / tmp1
	// 最高1 N % tmp1 + 1
	// 最高位first tmp1
	firstOneNum := 0
	if first == 1 {
		firstOneNum = n%tmp1 + 1
	} else {
		firstOneNum = tmp1
	}
	// 除去最高位之外，剩下1的数量
	// 最高位1 10(k-2次方) * (k-1) * 1
	// 最高位first 10(k-2次方) * (k-1) * first
	otherOneNum := first * (length - 1) * (tmp1 / 10)
	return firstOneNum + otherOneNum + countDigitOne(n%tmp1)
}

func getLenOfNum(num int) int {
	length := 0
	for num != 0 {
		length++
		num /= 10
	}
	return length
}

func powerBaseOf10(base int) int {
	return int(math.Pow10(base))
}
