package Go_LeetCode

func spiralOrder(matrix [][]int) []int {
	tR := 0
	tC := 0
	dR := len(matrix) - 1
	dC := len(matrix[0]) - 1
	ans := make([]int, 0)
	for tR <= dR && tC <= dC {
		printEdge(matrix, &ans, tR, tC, dR, dC)
		tR++
		tC++
		dR--
		dC--
	}
	return ans
}

func printEdge(m [][]int, ans *[]int, tR, tC, dR, dC int) {
	if tR == dR {
		for i := tC; i <= dC; i++ {
			*ans = append(*ans, m[tR][i])
		}
	} else if tC == dC {
		for i := tR; i <= dR; i++ {
			*ans = append(*ans, m[i][tC])
		}
	} else {
		curC := tC
		curR := tR
		for curC != dC {
			*ans = append(*ans, m[curR][curC])
			curC++
		}
		for curR != dR {
			*ans = append(*ans, m[curR][dC])
			curR++
		}
		for curC != tC {
			*ans = append(*ans, m[dR][curC])
			curC--
		}
		for curR != tR {
			*ans = append(*ans, m[curR][tC])
			curR--
		}
	}
	return
}
