package Go_LeetCode

type SegmentTree_GPT struct {
	MAXN   int
	arr    []int
	sum    []int
	lazy   []int
	change []int
	update []bool
}

func NewSegmentTree_GPT(origin []int) *SegmentTree {
	MAXN := len(origin) + 1
	arr := make([]int, MAXN)
	copy(arr[1:], origin)
	sum := make([]int, MAXN<<2)
	lazy := make([]int, MAXN<<2)
	change := make([]int, MAXN<<2)
	update := make([]bool, MAXN<<2)
	return &SegmentTree{MAXN, arr, sum, lazy, change, update}
}

func (st *SegmentTree_GPT) pushUp(rt int) {
	st.sum[rt] = st.sum[rt<<1] + st.sum[rt<<1|1]
}

func (st *SegmentTree_GPT) pushDown(rt, ln, rn int) {
	if st.update[rt] {
		st.update[rt<<1] = true
		st.update[rt<<1|1] = true
		st.change[rt<<1] = st.change[rt]
		st.change[rt<<1|1] = st.change[rt]
		st.lazy[rt<<1] = 0
		st.lazy[rt<<1|1] = 0
		st.sum[rt<<1] = st.change[rt] * ln
		st.sum[rt<<1|1] = st.change[rt] * rn
		st.update[rt] = false
	}
	if st.lazy[rt] != 0 {
		st.lazy[rt<<1] += st.lazy[rt]
		st.sum[rt<<1] += st.lazy[rt] * ln
		st.lazy[rt<<1|1] += st.lazy[rt]
		st.sum[rt<<1|1] += st.lazy[rt] * rn
		st.lazy[rt] = 0
	}
}

func (st *SegmentTree_GPT) Build(l, r, rt int) {
	if l == r {
		st.sum[rt] = st.arr[l]
		return
	}
	mid := (l + r) >> 1
	st.Build(l, mid, rt<<1)
	st.Build(mid+1, r, rt<<1|1)
	st.pushUp(rt)
}

func (st *SegmentTree_GPT) Update(L, R, C, l, r, rt int) {
	if L <= l && r <= R {
		st.update[rt] = true
		st.change[rt] = C
		st.sum[rt] = C * (r - l + 1)
		st.lazy[rt] = 0
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

func (st *SegmentTree_GPT) Add(L, R, C, l, r, rt int) {
	if L <= l && r <= R {
		st.sum[rt] += C * (r - l + 1)
		st.lazy[rt] += C
		return
	}
	mid := (l + r) >> 1
	st.pushDown(rt, mid-l+1, r-mid)
	if L <= mid {
		st.Add(L, R, C, l, mid, rt<<1)
	}
	if R > mid {
		st.Add(L, R, C, mid+1, r, rt<<1|1)
	}
	st.pushUp(rt)
}

func (st *SegmentTree_GPT) Query(L, R, l, r, rt int) int {
	if L <= l && r <= R {
		return st.sum[rt]
	}
	mid := (l + r) >> 1
	st.pushDown(rt, mid-l+1, r-mid)
	ans := 0
	if L <= mid {
		ans += st.Query(L, R, l, mid, rt<<1)
	}
	if R > mid {
		ans += st.Query(L, R, mid+1, r, rt<<1|1)
	}
	return ans
}
