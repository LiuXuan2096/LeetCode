package Go_LeetCode

// 时间复杂度O(N)，额外空间复杂度O(N)
func candy(ratings []int) int {
	if ratings == nil || len(ratings) == 0 {
		return 0
	}
	N := len(ratings)
	left := make([]int, N)
	for i := 1; i < N; i++ {
		if ratings[i-1] < ratings[i] {
			left[i] = left[i-1] + 1
		}
	}
	right := make([]int, N)
	for i := N - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		}
	}
	ans := 0
	for i := 0; i < N; i++ {
		ans += max_135(left[i], right[i])
	}
	return ans + N
}

func max_135(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
