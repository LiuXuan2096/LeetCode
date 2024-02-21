package Go_LeetCode

func twoSum(nums []int, target int) []int {
	// map的key是某个之前的数， value是这个数出现的位置
	numMap := make(map[int]int)
	for i, num := range nums {
		if index, ok := numMap[target-num]; ok {
			return []int{index, i}
		}
		numMap[num] = i
	}
	return []int{-1, -1}
}
