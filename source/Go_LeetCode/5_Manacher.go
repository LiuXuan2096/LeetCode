package Go_LeetCode

import (
	"math"
)

func LlongestPalindrome(s string) string {
	// "12132" -> "#1#2#1#3#2#"
	str := manacherString(s)
	var start, end, ansLength int
	var pArr = make([]int, len(str)) // 存放每个位置对应的回文半径的大小
	var C = -1
	// 算法描述中：R代表最右的扩成功的位置
	// 实际编码时：R表示最右的扩成功位置的，再下一个位置，即第一个范围扩充失败的位置
	var R = -1
	var max = math.MinInt32
	for i := 0; i < len(str); i++ {
		if i < R {
			// i落在R的范围内的情况
			pArr[i] = int(math.Min(float64(pArr[2*C-i]), float64(R-i)))
		} else {
			pArr[i] = 1
		}
		for i+pArr[i] < len(str) && i-pArr[i] > -1 {
			if str[i+pArr[i]] == str[i-pArr[i]] {
				pArr[i]++
			} else {
				break
			}
		}
		if pArr[i]+i > R {
			R = i + pArr[i]
			C = i
		}
		if max < pArr[i] {
			max = pArr[i]
			ansLength = max - 1
			// symPointer = i
			start = i - ansLength
			end = i + ansLength
		}

	}

	var ans string
	for i := start; i <= end; i++ {
		if str[i] != '#' {
			ans += string(str[i])
		}
	}

	//fmt.Println("str: ", string(str))
	//fmt.Println("对称字符为:", string(str[symPointer]))
	//fmt.Println("回文串长度为: ", ansLength)
	return string(ans)
}

func manacherString(str string) []byte {
	charArr := []byte(str)
	res := make([]byte, 2*len(str)+1)
	var index = 0
	for i := 0; i < len(res); i++ {
		if i&1 == 0 {
			res[i] = '#'
		} else {
			res[i] = charArr[index]
			index++
		}

	}
	return res
}
