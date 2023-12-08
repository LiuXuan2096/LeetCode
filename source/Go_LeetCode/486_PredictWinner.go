package Go_LeetCode

import "math"

func predictTheWinner(nums []int) bool {
	var N = len(nums)
	fmap := make([][]int, N) // 先手表
	for i := range fmap {
		fmap[i] = make([]int, N)
	}
	gmap := make([][]int, N) // 后手表
	for i := range gmap {
		gmap[i] = make([]int, N)
	}
	for i := 0; i < N; i++ {
		fmap[i][i] = nums[i]
	}
	for startCol := 1; startCol < N; startCol++ {
		var L = 0
		var R = startCol
		for R < N {
			fmap[L][R] = int(math.Max(float64(nums[L]+gmap[L+1][R]), float64(nums[R]+gmap[L][R-1])))
			gmap[L][R] = int(math.Min(float64(fmap[L+1][R]), float64(fmap[L][R-1])))
			L++
			R++
		}
	}

	if fmap[0][N-1] >= gmap[0][N-1] {
		return true
	} else {
		return false
	}
}
