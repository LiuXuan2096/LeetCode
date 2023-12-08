package Go_LeetCode

import "math"

//func coinChange(coins []int, amount int) int {
//	var ans = process7(coins, 0, amount)
//	if ans == math.MaxInt32 {
//		return -1
//	} else {
//		return ans
//	}
//
//}
//
//// arr[index...]面值，每种面值张数自由选择，
//// 搞出rest正好这么多钱，返回最小张数
//// 拿Integer.MAX_VALUE标记怎么都搞定不了
//func process7(arr []int, index int, rest int) int {
//	if index == len(arr) {
//		if rest == 0 {
//			return 0
//		} else {
//			return math.MaxInt32
//		}
//	} else {
//		var ans = math.MaxInt32
//		for zhang := 0; zhang*arr[index] <= rest; zhang++ {
//			var next = process7(arr, index+1, rest-zhang*arr[index])
//			if next != math.MaxInt32 {
//				ans = int(math.Min(float64(ans), float64(zhang+next)))
//			}
//		}
//		return ans
//	}
//}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	var N = len(coins)
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}
	dp[N][0] = 0
	for j := 1; j <= amount; j++ {
		dp[N][j] = math.MaxInt32
	}
	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest <= amount; rest++ {
			dp[index][rest] = dp[index+1][rest]
			if rest-coins[index] >= 0 && dp[index][rest-coins[index]] != math.MaxInt32 {
				dp[index][rest] = int(math.Min(float64(dp[index][rest]), float64(dp[index][rest-coins[index]]+1)))
			}
		}
	}
	if dp[0][amount] == math.MaxInt32 {
		return -1
	} else {
		return dp[0][amount]
	}

}
