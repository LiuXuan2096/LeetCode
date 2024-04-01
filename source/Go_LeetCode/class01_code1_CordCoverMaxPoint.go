package Go_LeetCode

import "math"

/*
给定一个有序数组arr，代表坐落在X轴上的点，给定一个正数K，
代表绳子的长度，返回绳子最多压中几个点？
即使绳子边缘处盖住点也算盖住
*/

func MaxPoints(arr []int, L int) int {
	left := 0
	right := 0
	N := len(arr)
	max := 0
	for left < N {
		for right < N && arr[right]-arr[left] <= L {
			right++
		}
		max = int(math.Max(float64(max), float64(right-left)))
		left++
	}
	return max
}
