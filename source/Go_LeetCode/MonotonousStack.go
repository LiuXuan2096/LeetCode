package Go_LeetCode

import "container/list"

func GetNearLess(arr []int) [][]int {
	var res = make([][]int, len(arr))
	for i := range res {
		res[i] = make([]int, 2)
	}
	var stack = make([]list.List, 10)
	for i := 0; i < len(arr); i++ {
		// i -> arr[i]进栈
		for len(stack) > 0 && arr[i] > arr[stack[len(stack)-1].Front().Value.(int)] {
			popls := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			var leftLessIndex int
			if len(stack) == 0 {
				leftLessIndex = -1
			} else {
				leftLessIndex = stack[len(stack)-1].Back().Value.(int)
			}
			for j := popls.Front(); j != nil; j = j.Next() {
				res[j.Value.(int)][0] = leftLessIndex
				res[j.Value.(int)][1] = i
			}
		}
		if len(stack) > 0 && arr[stack[len(stack)-1].Front().Value.(int)] == arr[i] {
			stack[len(stack)-1].PushBack(i)
		} else {
			var tempList list.List
			tempList.PushBack(i)
			stack = append(stack, tempList)
		}

	}
	for len(stack) > 0 {
		popls := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		var leftLessIndex int
		if len(stack) == 0 {
			leftLessIndex = -1
		} else {
			leftLessIndex = stack[len(stack)-1].Back().Value.(int)
		}
		for i := popls.Front(); i != nil; i = i.Next() {
			res[i.Value.(int)][0] = leftLessIndex
			res[i.Value.(int)][1] = -1
		}
	}
	return res
}
