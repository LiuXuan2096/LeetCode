package Go_LeetCode

func rotate(matrix [][]int) {
	a, b := 0, 0                            // 左上角顶点的坐标
	c, d := len(matrix)-1, len(matrix[0])-1 //右下角顶点的坐标
	for a < c {
		rotateEdge(matrix, a, b, c, d)
		a++
		b++
		c--
		d--
	}
}

func rotateEdge(m [][]int, a, b, c, d int) {
	for i := 0; i < d-b; i++ {
		tmp := m[a][b+i]
		m[a][b+i] = m[c-i][b]
		m[c-i][b] = m[c][d-i]
		m[c][d-i] = m[a+i][d]
		m[a+i][d] = tmp
	}
}
