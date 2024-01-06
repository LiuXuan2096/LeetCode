package Go_LeetCode

func MaxSumMinProduct(nums []int) int {
	var size = len(nums)
	var sum = make([]int, size)
	sum[0] = nums[0] // 前缀和数组
	for i := 1; i < size; i++ {
		sum[i] = sum[i-1] + nums[i]
	}

	nums = append(nums, 0)
	stack := make([]int, 0)
	var max int = 0
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] > nums[i] {
			left := 0
			if len(stack) > 1 {
				left = sum[stack[len(stack)-2]]
			}
			total := (sum[i-1] - left) * nums[stack[len(stack)-1]]
			if total > max {
				max = total
			}
			stack = stack[:len(stack)-1] // 相当于栈顶元素出栈
		}
		stack = append(stack, i)
	}
	return max % 1000000007

	//var stack list.List
	//for i := 0; i < size; i++ {
	//	for stack.Len() > 0 && nums[i] <= nums[stack.Back().Value.(int)] {
	//		var j = stack.Back().Value.(int)
	//		stack.Remove(stack.Back())
	//		var temp int
	//		if stack.Len() == 0 {
	//			temp = sum[i-1]
	//		} else {
	//			temp = (sum[i-1] - sum[stack.Back().Value.(int)]) * nums[j]
	//		}
	//		max = int64(int(math.Max(float64(temp), float64(max))))
	//	}
	//	stack.PushBack(i)
	//}
	//for stack.Len() > 0 {
	//	var j = stack.Back().Value.(int)
	//	stack.Remove(stack.Back())
	//	var temp int
	//	if stack.Len() == 0 {
	//		temp = sum[size-1]
	//	} else {
	//		temp = (sum[size-1] - sum[stack.Back().Value.(int)]) * nums[j]
	//	}
	//	max = int64(int(math.Max(float64(temp), float64(max))))
	//}
	//return (max % (10e9 + 7))
}

func mMaxSumMinProduct(nums []int) int {
	sum := make([]int, len(nums))
	sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i]
	}

	nums = append(nums, 0)
	stack := make([]int, 0)
	ans := 0
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] > nums[i] {
			left := 0
			if len(stack) > 1 {
				left = sum[stack[len(stack)-2]]
			}
			tot := (sum[i-1] - left) * nums[stack[len(stack)-1]]
			if tot > ans {
				ans = tot
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans % 1000000007
}
