package Go_LeetCode

func largest1BorderedSquare(grid [][]int) int {
	r := len(grid)
	c := len(grid[0])
	right := make([][]int, r)
	down := make([][]int, r)
	for i := range right {
		right[i] = make([]int, c)
		down[i] = make([]int, c)
	}
	setBorderMap(grid, right, down)
	for size := min_1139(r, c); size != 0; size-- {
		if hasSizeOfBorder(size, right, down) {
			return size * size
		}
	}
	return 0
}

func setBorderMap(m [][]int, right, down [][]int) {
	r := len(m)
	c := len(m[0])
	if m[r-1][c-1] == 1 {
		right[r-1][c-1] = 1
		down[r-1][c-1] = 1
	}
	for i := r - 2; i != -1; i-- {
		if m[i][c-1] == 1 {
			right[i][c-1] = 1
			down[i][c-1] = down[i+1][c-1] + 1
		}
	}
	for i := c - 2; i != -1; i-- {
		if m[r-1][i] == 1 {
			right[r-1][i] = right[r-1][i+1] + 1
			down[r-1][i] = 1
		}
	}
	for i := r - 2; i != -1; i-- {
		for j := c - 2; j != -1; j-- {
			if m[i][j] == 1 {
				right[i][j] = right[i][j+1] + 1
				down[i][j] = down[i+1][j] + 1
			}
		}
	}
}

func hasSizeOfBorder(size int, right, down [][]int) bool {
	r := len(right)
	c := len(right[0])
	for i := 0; i != r-size+1; i++ {
		for j := 0; j != c-size+1; j++ {
			if right[i][j] >= size && down[i][j] >= size && right[i+size-1][j] >= size && down[i][j+size-1] >= size {
				return true
			}
		}
	}
	return false
}

func min_1139(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
