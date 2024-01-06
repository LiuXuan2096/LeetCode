package Go_LeetCode

func swap(arr *[]int, p1, p2 int) {
	var temp = (*arr)[p1]
	(*arr)[p1] = (*arr)[p2]
	(*arr)[p2] = temp
}

func partition(arr []int, L, R, pivot int) [2]int {
	var left = L - 1
	var right = R + 1
	var cur = L
	for cur < right {
		if arr[cur] > pivot {
			swap(&arr, left+1, cur)
			left++
			cur++
		} else if arr[cur] < pivot {
			swap(&arr, right-1, cur)
			right--
		} else {
			cur++
		}
	}
	return [2]int{left + 1, right - 1}
}

func FfindKthLargest(nums []int, k int) int {
	k = k - 1
	var L = 0
	var R = len(nums) - 1
	var pivot = 0
	var pivotRange [2]int
	for L < R {
		pivot = nums[L+((R-L+1)>>1)]
		pivotRange = partition(nums, L, R, pivot)
		if k < pivotRange[0] {
			R = pivotRange[0] - 1
		} else if k > pivotRange[1] {
			L = pivotRange[1] + 1
		} else {
			return pivot
		}
	}
	return nums[L]
}
