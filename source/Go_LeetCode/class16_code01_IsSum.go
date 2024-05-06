package Go_LeetCode

/*
给定一个有正、有负、有0的数组arr，
给定一个整数k，
返回arr的子集是否能累加出k
1）正常怎么做？
2）如果arr中的数值很大，但是arr的长度不大，怎么做？
*/

// IsSum3 arr中的值可能为正，可能为负，可能为0
// 自由选择arr中的数字，能不能累加得到sum
// 经典动态规划
func IsSum3(arr []int, sum int) bool {
	if sum == 0 {
		return true
	}
	if arr == nil || len(arr) == 0 {
		return false
	}

	minV := 0
	maxV := 0
	for _, num := range arr {
		if num < 0 {
			minV += num
		}
		if num > 0 {
			maxV += num
		}
	}

	if sum < minV || sum > maxV {
		return false
	}

	N := len(arr)
	dp := make([][]bool, N)
	// dp[i][j]
	//
	//  0   1   2   3  4    5   6    7 (实际的对应)
	// -7  -6  -5  -4  -3   -2  -1   0（想象中）
	//
	// dp[0][-min] -> dp[0][7] -> dp[0][0]
	for i := range dp {
		dp[i] = make([]bool, maxV-minV+1)
	}

	dp[0][0] = true
	dp[0][arr[0]-minV] = true
	for i := 1; i < N; i++ {
		for j := minV; j <= maxV; j++ {
			dp[i][j-minV] = dp[i-1][j-minV]
			next := j - minV - arr[i]
			if next >= 0 && next <= maxV-minV {
				dp[i][j-minV] = dp[i][j-minV] || dp[i-1][next]
			}
		}
	}

	return dp[N-1][sum-minV]
}

//2）如果arr中的数值很大，但是arr的长度不大，怎么做？

// IsSum4 arr中的值可能为正，可能为负，可能为0
// 自由选择arr中的数字，能不能累加得到sum
// 分治的方法
// 如果arr中的数值特别大，动态规划方法依然会很慢
// 此时如果arr的数字个数不算多(40以内)，哪怕其中的数值很大，分治的方法也将是最优解
func IsSum4(arr []int, sum int) bool {
	if sum == 0 {
		return true
	}
	if arr == nil || len(arr) == 0 {
		return false
	}
	if len(arr) == 1 {
		return arr[0] == sum
	}

	N := len(arr)
	mid := N / 2

	leftSum := make(map[int]bool)
	rightSum := make(map[int]bool)
	// 0...mid-1
	process_16_01(arr, 0, mid, 0, leftSum)
	// mid...N-1
	process_16_01(arr, mid, N, 0, rightSum)
	// 单独查看，只使用左部分，能不能搞出sum
	// 单独查看，只使用右部分，能不能搞出sum
	// 左+右，联合能不能搞出sum
	// 左部分搞出所有累加和的时候，包含左部分一个数也没有，这种情况的，leftsum表里，0
	for l := range leftSum {
		if _, ok := rightSum[sum-l]; ok {
			return true
		}
	}
	return false
}

// process_16_01 arr[0...i-1]决定已经做完了！形成的累加和是pre
// arr[i...end - 1] end(终止)  所有数字随意选择，
// arr[0...end-1]所有可能的累加和存到ans里去
func process_16_01(arr []int, i, end, pre int, ans map[int]bool) {
	if i == end {
		ans[pre] = true
	} else {
		process_16_01(arr, i+1, end, pre, ans)
		process_16_01(arr, i+1, end, pre+arr[i], ans)
	}
}
