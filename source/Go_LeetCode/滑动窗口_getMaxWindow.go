package Go_LeetCode

import "container/list"

// GetMaxWindow
/*
窗口内最大值或最小值更新结构的实现
假设一个固定大小为W的窗口，依次划过arr，
返回每一次滑出状况的最大值
例如，arr = [4,3,5,4,3,3,6,7], W = 3
返回：[5,5,5,4,6,7]
*/
func GetMaxWindow(arr []int, w int) []int {
	if arr == nil || w < 1 || len(arr) < w {
		return nil
	}
	qmax := list.New()
	res := make([]int, len(arr)-w+1)
	i := 0

	for R := 0; R < len(arr); R++ {
		for qmax.Len() != 0 && arr[qmax.Back().Value.(int)] <= arr[R] {
			qmax.Remove(qmax.Back())
		}
		qmax.PushBack(R)

		if qmax.Front().Value.(int) == R-w {
			qmax.Remove(qmax.Front())
		}

		if R >= w-1 {
			res[i] = arr[qmax.Front().Value.(int)]
			i++
		}
	}
	return res
}
