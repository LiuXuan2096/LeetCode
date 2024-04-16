package Go_LeetCode

import "math"

type Node_1707 struct {
	min   int
	nexts [2]*Node_1707
}

type NumTrie_1707 struct {
	head *Node_1707
}

func NewNumTrie_1707() *NumTrie_1707 {
	return &NumTrie_1707{head: &Node_1707{min: math.MaxInt32}}
}

func (t *NumTrie_1707) add(num int) {
	cur := t.head
	t.head.min = min_1707(t.head.min, num)
	for move := 30; move >= 0; move-- {
		path := (num >> move) & 1
		if cur.nexts[path] == nil {
			cur.nexts[path] = &Node_1707{min: math.MaxInt32}
		}
		cur = cur.nexts[path]
		cur.min = min_1707(cur.min, num)
	}
}

// 这个结构中，已经收集了一票数字
// 请返回哪个数字与X异或的结果最大，返回最大结果
// 但是，只有<=m的数字，可以被考虑
func (t *NumTrie_1707) maxXorWithBehindM(x, m int) int {
	if t.head.min > m {
		return -1
	}
	// 一定存在某个数可以和x结合
	cur := t.head
	ans := 0
	for move := 30; move >= 0; move-- {
		path := (x >> move) & 1
		// 期待遇到的东西
		best := path ^ 1
		// best变成了实际遇到的
		if cur.nexts[best] == nil || cur.nexts[best].min > m {
			best ^= 1
		}
		ans |= (path ^ best) << move
		cur = cur.nexts[best]
	}
	return ans
}

func min_1707(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func maximizeXor(nums []int, queries [][]int) []int {
	trie := NewNumTrie_1707()
	for _, num := range nums {
		trie.add(num)
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = trie.maxXorWithBehindM(q[0], q[1])
	}
	return ans
}
