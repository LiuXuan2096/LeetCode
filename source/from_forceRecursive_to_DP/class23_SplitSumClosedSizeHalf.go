package from_forceRecursive_to_DP

import "math"

/*
给定一个正数数组arr，请把arr中所有的数分成两个集合
如果arr长度为偶数，两个集合包含数的个数要一样多
如果arr长度为奇数，两个集合包含数的个数必须只差一个
请尽量让两个集合的累加和接近
返回最接近的情况下，较小集合的累加和
*/

func SplitSumClosedSizeHalf(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	sum := 0
	for _, num := range arr {
		sum += num
	}
	sum /= 2
	if len(arr)%2 == 0 {
		return process_SplitSumClosedSize(arr, 0, len(arr)/2, sum)
	} else {
		return int(math.Max(float64(process_SplitSumClosedSize(arr, 0, len(arr)/2, sum)),
			float64(process_SplitSumClosedSize(arr, 0, len(arr)/2+1, sum))))
	}
}

// arr[i....]自由选择，挑选的个数一定要是picks个，
//累加和<=rest, 离rest最近的返回
func process_SplitSumClosedSize(arr []int, i, picks, rest int) int {
	if i == len(arr) {
		if picks == 0 {
			return 0
		}
		return -1
	}
	p1 := process_SplitSumClosedSize(arr, i+1, picks, rest)
	p2 := -1
	next := -1
	if arr[i] <= rest {
		next = process_SplitSumClosedSize(arr, i+1, picks-1, rest-arr[i])
	}
	if next != -1 {
		p2 = arr[i] + next
	}
	return int(math.Max(float64(p1), float64(p2)))
}

func SplitSumClosedSize_dp(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	sum := 0
	for _, num := range arr {
		sum += num
	}
	sum /= 2
	N := len(arr)
	M := (N + 1) / 2
	dp := make([][][]int, N+1)
	for i := range dp {
		dp[i] = make([][]int, M+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, sum+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	for rest := 0; rest <= sum; rest++ {
		dp[N][0][rest] = 0
	}
	for i := N - 1; i >= 0; i-- {
		for picks := 0; picks <= M; picks++ {
			for rest := 0; rest <= sum; rest++ {
				p1 := dp[i+1][picks][rest]
				p2 := -1
				next := -1
				if picks-1 >= 0 && arr[i] <= rest {
					next = dp[i+1][picks-1][rest-arr[i]]
				}
				if next != -1 {
					p2 = arr[i] + next
				}
				dp[i][picks][rest] = int(math.Max(float64(p1), float64(p2)))
			}
		}
	}
	if len(arr)%2 == 0 {
		return dp[0][len(arr)/2][sum]
	}
	return int(math.Max(float64(dp[0][len(arr)/2][sum]), float64(dp[0][(len(arr)/2)+1][sum])))
}
