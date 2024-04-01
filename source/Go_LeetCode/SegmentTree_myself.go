package Go_LeetCode

type Segment_Tree struct {
	// arr[]为原序列的信息从0开始，但在arr里是从1开始的
	// sum[]模拟线段树维护区间和
	// lazy[]为累加和懒惰标记
	// change[]为更新的值
	// update[]为更新慵懒标记
	MAXN   int
	arr    []int
	sum    []int
	lazy   []int
	change []int
	update []bool
}

func NewSegment_Tree(origin []int) *Segment_Tree {
	MaxN := len(origin) + 1
	arr := make([]int, MaxN)
	copy(arr[1:], origin)
	sum := make([]int, MaxN<<2)
	lazy := make([]int, MaxN<<2)
	change := make([]int, MaxN<<2)
	update := make([]bool, MaxN<<2)
	return &Segment_Tree{MaxN, arr, sum, lazy, change, update}
}

func (st *Segment_Tree) pushUp(rt int) {
	st.sum[rt] = st.sum[rt<<1] + st.sum[rt<<1|1]
}

// 之前的，所有懒增加，和懒更新，从父范围，发给左右两个子范围
// 分发策略是什么
// ln表示左子树元素结点个数，rn表示右子树结点个数
func (st *Segment_Tree) pushDown(rt, ln, rn int) {
	if st.update[rt] {
		// 更新左子树
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
		st.lazy[rt<<1|1] += st.lazy[rt]
		st.sum[rt<<1] = st.sum[rt<<1] + st.lazy[rt]*ln
		st.sum[rt<<1|1] = st.sum[rt<<1|1] + st.lazy[rt]*rn
		st.lazy[rt] = 0
	}
}

// Build 在初始化阶段，先把sum数组，填好
// 在arr[l~r]范围上，去build，1~N，
// rt : 这个范围在sum中的下标
func (st *Segment_Tree) Build(l, r, rt int) {
	if l == r {
		st.sum[rt] = st.arr[l]
		return
	}
	mid := (l + r) >> 1
	st.Build(l, mid, rt<<1)
	st.Build(mid+1, r, rt<<1|1)
	st.pushUp(rt)
}

// Update L~R  所有的值变成C
// l~r  rt
// L-R表示要更新的区间的范围
// l-r表示线段树的某个节点对应的原数组中的对应区间的范围
func (st *Segment_Tree) Update(L, R, C, l, r, rt int) {
	if L <= l && r <= R {
		st.update[rt] = true
		st.change[rt] = C
		st.sum[rt] = C * (r - l + 1)
		st.lazy[rt] = 0
		return
	}
	// 当前任务躲不掉，无法懒更新，要往下发
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

// Add L~R, C 任务！
// rt，l~r
// L-R表示要更新的区间的范围
// l-r表示线段树的某个节点对应的原数组中的对应区间的范围
func (st *Segment_Tree) Add(L, R, C, l, r, rt int) {
	if L <= l && r <= R {
		st.lazy[rt] += C
		st.sum[rt] += C * (r - l + 1)
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

// Query 查询L~R,范围的累加和
// rt，l~r
// L-R表示要更新的区间的范围
// l-r表示线段树的某个节点对应的原数组中的对应区间的范围
func (st *Segment_Tree) Query(L, R, l, r, rt int) int {
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
