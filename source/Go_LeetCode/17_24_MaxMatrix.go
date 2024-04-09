package Go_LeetCode

import "math"

func getMaxMatrix(m [][]int) []int {
	N := len(m)
	M := len(m[0])
	max := math.MinInt32
	cur := 0
	a, b, c, d := 0, 0, 0, 0

	for i := 0; i < N; i++ {
		s := make([]int, M)
		for j := i; j < N; j++ {
			cur = 0
			begin := 0
			for k := 0; k < M; k++ {
				s[k] += m[j][k]
				cur += s[k]
				if max < cur {
					max = cur
					a = i
					b = begin
					c = j
					d = k
				}
				if cur < 0 {
					cur = 0
					begin = k + 1
				}
			}
		}
	}
	return []int{a, b, c, d}
}
