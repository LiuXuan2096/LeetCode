package Go_LeetCode

/*
目前这个版本的答案有问题，只能通过三个测试用例
*/
type NumMatrix struct {
	tree [][]int
	nums [][]int
	N    int
	M    int
}

func Constructor(matrix [][]int) NumMatrix {
	row := len(matrix)
	col := len(matrix[0])
	tree := make([][]int, row+1)
	nums := make([][]int, row)
	for i := range tree {
		tree[i] = make([]int, col+1)
	}
	for i := range nums {
		nums[i] = make([]int, col)
	}
	indexTree := NumMatrix{tree: tree, nums: nums, N: row, M: col}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			indexTree.Update(i, j, matrix[i][j])
		}
	}
	return indexTree

}

func (this *NumMatrix) Update(row int, col int, val int) {
	if this.N == 0 || this.M == 0 {
		return
	}
	add := val - this.nums[row][col]
	for i := row + 1; i <= this.N; i += (i & -i) {
		for j := col + 1; j <= this.M; j += (j & -j) {
			this.tree[i][j] += add
		}
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if this.N == 0 || this.M == 0 {
		return 0
	}
	return this.sum(row2, col2) + this.sum(row1-1, col1-1) - this.sum(row2, col1-1) - this.sum(row1-1, col2)
}

func (this *NumMatrix) sum(row, col int) int {
	sum := 0
	for i := row + 1; i > 0; i -= (i & (-i)) {
		for j := col + 1; j > 0; j -= (j & (-j)) {
			sum += this.tree[i][j]
		}
	}

	return sum
}
