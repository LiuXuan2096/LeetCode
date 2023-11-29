package Go_LeetCode

var parent []int
var size []int
var help []int
var row int
var col int
var sets int

func numIslands2(m int, n int, positions [][]int) []int {
	UinonFind1(m, n)
	var ans []int
	for _, position := range positions {
		ans = append(ans, connect(position[0], position[1]))
	}
	return ans
}

func UinonFind1(m, n int) {
	row = m
	col = n
	length := row * col
	sets = 0
	parent = make([]int, length)
	size = make([]int, length)
	help = make([]int, length)
}

func index(r, c int) int {
	return r*col + c
}

func find(i int) int {
	var helpIndex = 0
	for parent[i] != i {
		help[helpIndex] = i
		helpIndex++
		i = parent[i]
	}
	for helpIndex--; helpIndex >= 0; helpIndex-- {
		parent[help[helpIndex]] = i
	}
	return i
}

func union(r1, c1, r2, c2 int) {
	if r1 < 0 || r1 == row || r2 < 0 || r2 == row || c1 < 0 || c1 == col || c2 < 0 || c2 == col {
		return
	}
	var i1 = index(r1, c1)
	var i2 = index(r2, c2)
	if size[i1] == 0 || size[i2] == 0 {
		return
	}
	var f1 = find(i1)
	var f2 = find(i2)
	if f1 != f2 {
		if size[f1] > size[f2] {
			size[f1] += size[f2]
			parent[f2] = f1
		} else {
			size[f2] += size[f1]
			parent[f1] = f2
		}
		sets--
	}
}

func connect(r, c int) int {
	var index = index(r, c)
	if size[index] == 0 {
		parent[index] = index
		size[index] = 1
		sets++
		union(r-1, c, r, c)
		union(r+1, c, r, c)
		union(r, c-1, r, c)
		union(r, c+1, r, c)
	}
	return sets
}
