package Go_LeetCode

func minPatches(nums []int, n int) int {
	patches := 0  //缺多少个数字
	rangeVal := 0 //已经完成了1~rangeVal的小目标

	// 由于题目说明arr已经有序，这里不需要再排序
	// sort.Ints(arr) // 如果arr不是有序的，取消注释这行代码

	for _, num := range nums {
		// 要求：1 ~ num-1 范围被搞定！
		for num-1 > rangeVal {
			rangeVal += rangeVal + 1
			patches++
			if rangeVal >= n {
				return patches
			}
		}
		// 要求被满足了！
		rangeVal += num
		if rangeVal >= n {
			return patches
		}
	}
	// 如果没有数组或者数组里的数字不足以达到aim，继续添加
	for n > rangeVal {
		rangeVal += rangeVal + 1
		patches++
	}
	return patches
}
