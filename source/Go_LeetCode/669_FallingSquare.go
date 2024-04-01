package Go_LeetCode

func FallingSquares(positions [][]int) []int {
	// m := index_669(positions)
	N := len(positions)
	segmentTree := NewSegmentTree_669(N)
	// maxHeight := 0
	res := make([]int, 0)
	for _, info := range positions {
		x, h := info[0], info[1]
		cur := segmentTree.query(x, x+h-1, 1, N, 1)
		segmentTree.Update(x, x+h-1, cur+h, 1, N, 1)
		res = append(res, segmentTree.max[1])
	}
	return res
}

type SegmentTree_669 struct {
	max    []int
	change []int
	update []bool
}

func NewSegmentTree_669(size int) *SegmentTree_669 {
	N := size + 1
	return &SegmentTree_669{
		max:    make([]int, N<<2),
		change: make([]int, N<<2),
		update: make([]bool, N<<2),
	}
}

func (st *SegmentTree_669) pushUp(rt int) {
	st.max[rt] = max_669(st.max[rt<<1], st.max[rt<<1|1])
}

func (st *SegmentTree_669) pushDown(rt, ln, rn int) {
	if st.update[rt] {
		st.update[rt<<1] = true
		st.update[rt<<1|1] = true
		st.change[rt<<1] = st.change[rt]
		st.change[rt<<1|1] = st.change[rt]
		st.max[rt<<1] = st.change[rt]
		st.max[rt<<1|1] = st.change[rt]
		st.update[rt] = false
	}
}

func (st *SegmentTree_669) Update(L, R, C, l, r, rt int) {
	if L <= l && r <= R {
		st.update[rt] = true
		st.change[rt] = C
		st.max[rt] = C
		return
	}
	mid := (l + r) >> 1
	st.pushDown(rt, mid-l+1, r-mid)
	if L <= mid {
		st.Update(L, R, C, l, mid, rt<<1)
	}
	if R > mid {
		st.Update(L, R, C, mid+1, r, rt<<1|1)
	}
	st.pushUp(rt)
}

func (st *SegmentTree_669) query(L, R, l, r, rt int) int {
	if L <= l && r <= R {
		return st.max[rt]
	}
	mid := (l + r) >> 1
	st.pushDown(rt, mid-l+1, r-mid)
	left, right := 0, 0
	if L <= mid {
		left = st.query(L, R, l, mid, rt<<1)
	}
	if R > mid {
		right = st.query(L, R, mid+1, r, rt<<1|1)
	}
	return max(left, right)
}

func max_669(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//func index_669(positions [][]int) map[int]int {
//	pos := make([]int, 0)
//	for _, arr := range positions {
//		pos = append(pos, arr[0])
//		pos = append(pos, arr[0]+arr[1]-1)
//	}
//	fmt.Printf("排序前的pos: %v\n", pos)
//	sort.Ints(pos)
//	fmt.Printf("排序后的pos: %v\n", pos)
//	m := make(map[int]int)
//	count := 0
//	for _, index := range pos {
//		if _, ok := m[index]; !ok {
//			count++
//			m[index] = count
//		}
//	}
//	fmt.Printf("map: %v\n", m)
//	return m
//}
