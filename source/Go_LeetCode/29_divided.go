package Go_LeetCode

func divide(dividend int, divisor int) int {
	if divisor == -1 {
		if dividend == -(1 << 31) {
			return (1 << 31) - 1
		}
		return -dividend
	}

	if dividend == -(1 << 31) {
		if divisor == 1 {
			return -(1 << 31)
		}
		res := div(add(dividend, 1), divisor)
		return add(res, div(minus(dividend, multi(res, divisor)), divisor))
	}
	return div(dividend, divisor)
}

func add(a, b int) int {
	var sum int
	for b != 0 {
		sum = a ^ b      // 不进位相加的信息
		b = (a & b) << 1 //进位信息 b不为0 说明需要进位
		a = sum
	}
	// b==0 时，说明相加不需要进位，此时异或运算的结果就是
	// 两数相加的结果
	return sum
}

func negNum(n int) int {
	return add(^n, 1)
}

func minus(a, b int) int {
	return add(a, negNum(b))
}

func multi(a, b int) int {
	var res int
	for b != 0 {
		if (b & 1) != 0 {
			res = add(res, a)
		}
		a <<= 1
		b >>= 1
	}
	return res
}

func isNeg(n int) bool {
	return n < 0
}

func div(a, b int) int {
	x := a
	y := b
	var res int
	for i := 31; i > negNum(1); i = minus(i, 1) {
		if (x >> i) >= y {
			res |= (1 << i)
			x = minus(x, multi(y, 1<<i))
		}
	}
	if isNeg(a) != isNeg(b) {
		return negNum(res)
	}
	return res
}
