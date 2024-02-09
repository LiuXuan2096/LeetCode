package Go_LeetCode

type IndexTree struct {
	tree []int
	N    int
}

func NewIndexTree(size int) *IndexTree {
	N := size
	tree := make([]int, N+1)
	return &IndexTree{tree: tree, N: N}
}

func (it *IndexTree) Sum(index int) int {
	ret := 0
	for index > 0 {
		ret += it.tree[index]
		index -= index & -index // index & -index 提取出index最右侧的1来
	}
	return ret
}

func (it *IndexTree) Add(index, d int) {
	for index <= it.N {
		it.tree[index] += d
		index += index & -index
	}
}
