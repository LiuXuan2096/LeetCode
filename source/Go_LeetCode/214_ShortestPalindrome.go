package Go_LeetCode

import (
	"math"
)

func SshortestPalindrome(s string) string {
	str := manacherString2(s)
	var pArr = make([]int, len(str)) // 存放s中每个字符对应的回文半径
	var C = 50001
	var R = 50001
	var finalRadius = -1
	for i := len(str) - 1; i >= 0; i-- {
		if R < i {
			pArr[i] = int(math.Min(float64(pArr[2*C-i]), float64(i-R)))
		} else {
			pArr[i] = 1
		}
		for i-pArr[i] >= 0 && i+pArr[i] < len(str) {
			if str[i-pArr[i]] == str[i+pArr[i]] {
				pArr[i]++
			} else {
				break
			}
		}
		if i-pArr[i] < R {
			R = i - pArr[i]
			C = i
		}
		if R == -1 {
			finalRadius = pArr[i]
			break
		}
	}
	var end = C + finalRadius
	var start = len(str) - 1
	res := make([]byte, len(str)+(len(str)-end))
	for j := 0; start >= end; j++ {
		res[j] = str[start]
		start--
	}
	for i := range str {
		res = append(res, str[i])
	}
	ans := make([]byte, 0)
	for i := 0; i < len(res); i++ {
		if res[i] != '#' {
			ans = append(ans, res[i])
		}
	}
	//fmt.Println("start: ", start)
	//fmt.Println("对称点: ", string(str[C]))
	//fmt.Println("finalRadius: ", finalRadius)
	//fmt.Println("res: ", string(res))

	//answer := make([]rune, 0, len(ans))
	//for _, v := range ans {
	//	if v == 0 {
	//		continue
	//	}
	//	answer = append(answer, v)
	//}
	//return string(answer)
	answer := ""
	for _, v := range ans {
		if v == 0 {
			continue
		}
		answer += string(v)
	}
	return answer
	//return string(ans)
}

func manacherString2(str string) []byte {
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
