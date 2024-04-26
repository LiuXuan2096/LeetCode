package Go_LeetCode

// 不是最优解 常数时间还可以优化，但是这种解法容易理解
func longestConsecutive(nums []int) int {
	headMap := make(map[int]int)
	tailMap := make(map[int]int)
	visited := make(map[int]bool)

	for _, num := range nums {
		if !visited[num] {
			visited[num] = true
			headMap[num] = 1
			tailMap[num] = 1

			if tailMap[num-1] != 0 {
				preLen := tailMap[num-1]
				preHead := num - preLen
				headMap[preHead] = preLen + 1
				tailMap[num] = preLen + 1
				delete(headMap, num)
				delete(tailMap, num-1)
			}

			if headMap[num+1] != 0 {
				preLen := tailMap[num]
				preHead := num - preLen + 1
				postLen := headMap[num+1]
				postTail := num + postLen
				headMap[preHead] = preLen + postLen
				tailMap[postTail] = preLen + postLen
				delete(headMap, num+1)
				delete(tailMap, num)
			}
		}
	}

	ans := 0
	for _, length := range headMap {
		if length > ans {
			ans = length
		}
	}

	return ans
}
