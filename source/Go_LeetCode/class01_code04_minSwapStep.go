package Go_LeetCode

import "math"

/*
MinSteps1
一个数组中只有两种字符'G'和'B'，
可以让所有的G都放在左侧，所有的B都放在右侧
或者可以让所有的G都放在右侧，所有的B都放在左侧，
但是只能在相邻字符之间进行交换操作，返回至少需要交换几次
*/
func MinSteps1(s string) int {
	if s == "" {
		return 0
	}
	str := []rune(s)
	step1 := 0
	gi := 0
	for i := 0; i < len(str); i++ {
		if str[i] == 'G' {
			step1 += i - gi
			gi++
		}
	}
	step2 := 0
	bi := 0
	for i := 0; i < len(str); i++ {
		if str[i] == 'B' {
			step2 += i - bi
			bi++
		}
	}
	return int(math.Min(float64(step1), float64(step2)))
}
