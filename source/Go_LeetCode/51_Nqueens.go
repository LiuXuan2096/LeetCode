package Go_LeetCode

import (
	"fmt"
	"math"
)

func SolveNQueens(n int) [][]string {
	var cases = make([][]int, 0)
	//for i := range cases {
	//	cases[i] = make([]int, n)
	//}
	var res = make([]int, n)
	var counts = process8(0, &cases, n, res)
	// fmt.Println(cases)
	// fmt.Println("方法数: ", counts)
	if counts == 0 {
		return nil
	} else {
		var ans = make([][]string, counts)
		for i := range ans {
			ans[i] = make([]string, n)
		}
		for i := 0; i < counts; i++ {
			for j := 0; j < n; j++ {
				var fillStr = ""
				for k := 0; k < n; k++ {
					if k == cases[i][j] {
						fillStr += "Q"
					} else {
						fillStr += "."
					}
				}
				fmt.Println("filstr: ", fillStr)
				ans[i][j] = fillStr
			}
		}
		return ans
	}
}

func process8(i int, cases *[][]int, n int, res []int) int {
	if i == n {
		// fmt.Println("record: ", res)
		var temp = make([]int, n)
		for i := 0; i < n; i++ {
			temp[i] = res[i]
		}
		(*cases) = append(*cases, temp)
		return 1

	}
	var counts = 0
	for j := 0; j < n; j++ {
		if isValid(res, i, j) {
			res[i] = j
			counts += process8(i+1, cases, n, res)
		}
	}
	return counts
}

func isValid(cases []int, i, j int) bool {
	for k := 0; k < i; k++ {
		if j == cases[k] || math.Abs(float64(cases[k]-j)) == math.Abs(float64(k-i)) {
			return false
		}
	}
	return true
}
