package Go_LeetCode

import "container/list"

// 这个单调栈实现的不对
func GetNearLess2(arr []int) [][]int {
	res := make([][]int, len(arr))
	stack := list.New()

	for i := 0; i < len(arr); i++ {
		for stack.Len() > 0 && arr[stack.Back().Value.(int)] > arr[i] {
			popIs := stack.Remove(stack.Back()).([]int)
			leftLessIndex := -1
			if stack.Len() > 0 {
				leftLessIndex = stack.Back().Value.(int)
			}
			for _, popi := range popIs {
				res[popi][0] = leftLessIndex
				res[popi][1] = i
			}
		}
		if stack.Len() > 0 && arr[stack.Back().Value.(int)] == arr[i] {
			stack.Back().Value = append(stack.Back().Value.([]int), i)
		} else {
			stack.PushBack([]int{i})
		}
	}

	for stack.Len() > 0 {
		popIs := stack.Remove(stack.Back()).([]int)
		leftLessIndex := -1
		if stack.Len() > 0 {
			leftLessIndex = stack.Back().Value.(int)
		}
		for _, popi := range popIs {
			res[popi][0] = leftLessIndex
			res[popi][1] = -1
		}
	}
	return res
}
