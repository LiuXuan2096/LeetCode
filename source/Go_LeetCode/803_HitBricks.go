package Go_LeetCode

func hitBricks(grid [][]int, hits [][]int) []int {
	for i := 0; i < len(hits); i++ {
		if grid[hits[i][0]][hits[i][1]] == 1 {
			grid[hits[i][0]][hits[i][1]] = 2
		}
	}
	uf := newUnionFind(grid)
	ans := make([]int, len(hits))
	for i := len(hits) - 1; i >= 0; i-- {
		if grid[hits[i][0]][hits[i][1]] == 2 {
			ans[i] = uf.Finger(hits[i][0], hits[i][1])
		}
	}
	return ans
}

// 手撸一个并查集

type UnionFind_803 struct {
	N, M       int
	CeilingAll int     // 有多少块砖，连到了天花板上
	Grid       [][]int // 原始矩阵，因为炮弹的影响，1 -> 2
	// cellingSet[i] = true; i 是头节点，
	//所在的集合是天花板集合,和天花板相连的集合
	CeilingSet []bool
	FatherMap  []int
	SizeMap    []int
	Stack      []int
}

func newUnionFind(matrix [][]int) *UnionFind_803 {
	uf := &UnionFind_803{}
	uf.InitSpace(matrix)
	uf.InitConnect()
	return uf
}

func (uf *UnionFind_803) InitSpace(matrix [][]int) {
	uf.Grid = matrix
	uf.N = len(uf.Grid)
	uf.M = len(uf.Grid[0])
	all := uf.N * uf.M
	uf.CeilingSet = make([]bool, all)
	uf.FatherMap = make([]int, all)
	uf.SizeMap = make([]int, all)
	uf.Stack = make([]int, all)
	for row := 0; row < uf.N; row++ {
		for col := 0; col < uf.M; col++ {
			if uf.Grid[row][col] == 1 {
				index := row*uf.M + col
				uf.FatherMap[index] = index
				uf.SizeMap[index] = 1
				if row == 0 {
					uf.CeilingSet[index] = true
					uf.CeilingAll++
				}
			}
		}
	}
}

func (uf *UnionFind_803) InitConnect() {
	for row := 0; row < uf.N; row++ {
		for col := 0; col < uf.M; col++ {
			uf.Union(row, col, row-1, col)
			uf.Union(row, col, row+1, col)
			uf.Union(row, col, row, col-1)
			uf.Union(row, col, row, col+1)
		}
	}
}

// Find 返回某个节点所属集合的代表节点的索引
func (uf *UnionFind_803) Find(row, col int) int {
	stackSize := 0
	index := row*uf.M + col
	for index != uf.FatherMap[index] {
		uf.Stack[stackSize] = index
		stackSize++
		index = uf.FatherMap[index]
	}
	for stackSize != 0 {
		stackSize--
		uf.FatherMap[uf.Stack[stackSize]] = index
	}
	return index
}

func (uf *UnionFind_803) Union(r1, c1, r2, c2 int) {
	if uf.Valid(r1, c1) && uf.Valid(r2, c2) {
		father1 := uf.Find(r1, c1)
		father2 := uf.Find(r2, c2)
		if father1 != father2 {
			size1 := uf.SizeMap[father1]
			size2 := uf.SizeMap[father2]
			status1 := uf.CeilingSet[father1]
			status2 := uf.CeilingSet[father2]
			if size1 <= size2 {
				uf.FatherMap[father1] = father2
				uf.SizeMap[father2] = size1 + size2
				if status1 != status2 {
					uf.CeilingSet[father2] = true
					if status1 {
						uf.CeilingAll += size2
					} else {
						uf.CeilingAll += size1
					}
				}
			} else {
				uf.FatherMap[father2] = father1
				uf.SizeMap[father1] = size1 + size2
				if status1 != status2 {
					uf.CeilingSet[father1] = true
					if status1 {
						uf.CeilingAll += size2
					} else {
						uf.CeilingAll += size1
					}
				}
			}
		}
	}
}

func (uf *UnionFind_803) Valid(row, col int) bool {
	return row >= 0 && row < uf.N && col >= 0 && col < uf.M && uf.Grid[row][col] == 1
}

func (uf *UnionFind_803) CeilingNum() int {
	return uf.CeilingAll
}

// Finger 模拟炮弹炸掉砖块的逻辑
func (uf *UnionFind_803) Finger(row, col int) int {
	uf.Grid[row][col] = 1
	cur := row*uf.M + col
	if row == 0 {
		uf.CeilingSet[cur] = true
		uf.CeilingAll++
	}
	uf.FatherMap[cur] = cur
	uf.SizeMap[cur] = 1
	pre := uf.CeilingAll
	uf.Union(row, col, row-1, col)
	uf.Union(row, col, row+1, col)
	uf.Union(row, col, row, col-1)
	uf.Union(row, col, row, col+1)
	now := uf.CeilingAll
	if row == 0 {
		return now - pre
	} else {
		if now == pre {
			return 0
		} else {
			return now - pre - 1
		}
	}
}
