package Go_LeetCode

import (
	"math"
)

/*
当前版本的解答有错误
*/

func MinimumDifference(nums []int) int {
	var sum = 0
	var max = 0
	for i := range nums {
		sum += nums[i]
		max = int(math.Max(float64(max), float64(nums[i])))
	}
	var newNums = make([]int, len(nums))
	for i := range newNums {
		newNums[i] = nums[i] + max
	}
	var newSum = sum + max*len(nums)
	var sum1 = process7(newNums, 0, len(nums)>>1, newSum>>1)
	return int(math.Abs(float64(newSum - sum1 - sum1)))
}

// arr[i....]自由选择，挑选的个数一定要是picks个，累加和<=rest, 离rest最近的返回
func process7(nums []int, i, picks, rest int) int {
	if len(nums) == i {
		if picks == 0 {
			return 0
		} else {
			return -1
		}
	}
	var p1 = process7(nums, i+1, picks, rest)
	var p2 = -1
	var next = -1
	if nums[i] <= rest {
		next = process7(nums, i+1, picks-1, rest-nums[i])
	}
	if next != -1 {
		p2 = nums[i] + next
	}
	return int(math.Max(float64(p1), float64(p2)))
}

//func process7(arr []int, i, picks, rest int) int {
//	if i == len(arr) {
//		if picks == 0 {
//			return 0
//		} else {
//			return math.MinInt32
//		}
//	} else {
//		var p1 = process7(arr, i+1, picks, rest) // 没有使用i位置的数
//		var p2 = math.MinInt32                   // 挑选了i位置的数
//		var next = math.MinInt32
//		//if arr[i] <= rest {
//		//	next = process7(arr, i+1, picks-1, rest-arr[i])
//		//}
//		next = process7(arr, i+1, picks-1, rest-arr[i])
//		if next != math.MinInt32 {
//			p2 = arr[i] + next
//		}
//		return int(math.Max(float64(p1), float64(p2)))
//	}
//}

func MminimumDifference(nums []int) int {
	var sum = 0
	var max = 0
	for i := range nums {
		sum += nums[i]
		max = int(math.Max(float64(max), float64(nums[i])))
	}
	var newNums = make([]int, len(nums))
	for i := range newNums {
		newNums[i] = nums[i] + max
	}
	var newSum = sum + max*len(nums)
	var halfSum = newSum >> 1
	var N = len(nums)
	var picks = N >> 1
	dp := make([][][]int, N+1)
	for i := range dp {
		dp[i] = make([][]int, picks+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, halfSum+1)
		}
	}
	// 初始化动态规划表
	for i := 0; i <= N; i++ {
		for j := 0; j <= picks; j++ {
			for k := 0; k <= halfSum; k++ {
				dp[i][j][k] = -1
			}
		}
	}
	for rest := 0; rest <= halfSum; rest++ {
		dp[N][0][rest] = 0
	}
	for i := N - 1; i >= 0; i-- {
		for pick := 0; pick <= picks; pick++ {
			for rest := 0; rest <= halfSum; rest++ {
				var p1 = dp[i+1][pick][rest]
				var p2 = -1
				var next = -1
				if pick-1 >= 0 && nums[i] <= rest {
					next = dp[i+1][pick-1][rest-nums[i]]
				}
				if next != -1 {
					p2 = nums[i] + next
				}
				dp[i][pick][rest] = int(math.Max(float64(p1), float64(p2)))
			}

		}
	}
	return int(math.Abs(float64(newSum - dp[0][picks][halfSum] - dp[0][picks][halfSum])))
}
