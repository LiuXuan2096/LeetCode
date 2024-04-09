package Go_LeetCode

import "sort"

/*
 * 今日头条原题
 *
 * 数组为{3, 2, 2, 3, 1}，查询为(0, 3, 2)。意思是在数组里下标0~3这个范围上，有几个2？返回2。
 * 假设给你一个数组arr，对这个数组的查询非常频繁，请返回所有查询的结果
 *
 */
type QueryBox struct {
	mapPositions map[int][]int
}

func NewQueryBox(arr []int) *QueryBox {
	q := &QueryBox{
		mapPositions: make(map[int][]int),
	}
	for i, v := range arr {
		q.mapPositions[v] = append(q.mapPositions[v], i)
	}
	return q
}

func (q *QueryBox) Query(L, R, value int) int {
	indexArr, ok := q.mapPositions[value]
	if !ok {
		return 0
	}
	// 查询 < L 的下标有几个
	a := q.countLess(indexArr, L)
	// 查询 < R+1 的下标有几个
	b := q.countLess(indexArr, R+1)
	return b - a
}

// 在有序数组arr中，用二分的方法数出<limit的数有几个
// 也就是用二分法，找到<limit的数中最右的位置
func (q *QueryBox) countLess(arr []int, limit int) int {
	mostRight := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= limit
	})
	return mostRight
}
