package Go_LeetCode

func numIslands(grid [][]byte) int {
	var islands = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				islands++
				infect(grid, i, j)
			}
		}
	}
	return islands
}

// 从i j位置出发，把所有练成一片的1字符变成0
func infect(grid [][]byte, i, j int) {
	if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) || grid[i][j] != '1' {
		return
	}
	grid[i][j] = 0
	infect(grid, i-1, j)
	infect(grid, i, j+1)
	infect(grid, i+1, j)
	infect(grid, i, j-1)
}
