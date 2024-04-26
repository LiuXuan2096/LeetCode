package Go_LeetCode

/*
https://leetcode-cn.com/problems/boolean-evaluation-lcci/
*/

type Info_0814 struct {
	t int //值为true的方法有多少种
	f int //值为false的方法有多少种
}

// 限制:
// L...R上，一定有奇数个字符
// L位置的字符和R位置的字符，非0即1，不能是逻辑符号！
// 返回str[L...R]这一段，为true的方法数，和false的方法数
func eval(str []rune, L, R int, dp [][]*Info_0814) *Info_0814 {
	if dp[L][R] != nil {
		return dp[L][R]
	}
	t := 0
	f := 0
	if L == R {
		if str[L] == '1' {
			t = 1
		} else {
			f = 1
		}
	} else {
		// L..R >=3
		// 每一个种逻辑符号，split枚举的东西
		// 都去试试最后结合
		for split := L + 1; split < R; split += 2 {
			leftInfo := eval(str, L, split-1, dp)
			rightInfo := eval(str, split+1, R, dp)
			a, b, c, d := leftInfo.t, leftInfo.f, rightInfo.t, rightInfo.f
			switch str[split] {
			case '&':
				t += a * c
				f += b*c + b*d + a*d
			case '|':
				t += a*c + a*d + b*c
				f += b * d
			case '^':
				t += a*d + b*c
				f += a*c + b*d

			}
		}
	}
	dp[L][R] = &Info_0814{t, f}
	return dp[L][R]
}

func countEval(s string, result int) int {
	if s == "" {
		return 0
	}
	exp := []rune(s)
	N := len(exp)
	dp := make([][]*Info_0814, N)
	for i := range dp {
		dp[i] = make([]*Info_0814, N)
	}
	allInfo := eval(exp, 0, N-1, dp)
	if result == 1 {
		return allInfo.t
	}
	return allInfo.f
}
