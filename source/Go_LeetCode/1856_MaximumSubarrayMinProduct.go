package Go_LeetCode

import (
	"container/list"
	"math"
)

/*
这个版本的代码有问题，没通过全部测试用例
*/
func MaxSumMinProduct(nums []int) int64 {
	var size = len(nums)
	var preSum = make([]int, size)
	preSum[0] = nums[0] // 前缀和数组
	for i := 1; i < size; i++ {
		preSum[i] = preSum[i-1] + nums[i]
	}
	var max int64 = math.MinInt32
	var stack list.List
	for i := 0; i < size; i++ {
		for stack.Len() > 0 && nums[i] <= nums[stack.Back().Value.(int)] {
			var j = stack.Back().Value.(int)
			stack.Remove(stack.Back())
			var temp int
			if stack.Len() == 0 {
				temp = preSum[i-1]
			} else {
				temp = (preSum[i-1] - preSum[stack.Back().Value.(int)]) * nums[j]
			}
			max = int64(int(math.Max(float64(temp), float64(max))))
		}
		stack.PushBack(i)
	}
	for stack.Len() > 0 {
		var j = stack.Back().Value.(int)
		stack.Remove(stack.Back())
		var temp int
		if stack.Len() == 0 {
			temp = preSum[size-1]
		} else {
			temp = (preSum[size-1] - preSum[stack.Back().Value.(int)]) * nums[j]
		}
		max = int64(int(math.Max(float64(temp), float64(max))))
	}
	return (max % (10e9 + 7))
}
