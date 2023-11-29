package Go_LeetCode

import Go_LeetCode "LeetCode-go/source/Go_Datastructure_algorithm"

func FindCircleNum(isConnected [][]int) int {
	var N = len(isConnected)
	Go_LeetCode.InitUnionFind(N)
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if isConnected[i][j] == 1 {
				Go_LeetCode.Union(i, j)
			}
		}
	}

	return Go_LeetCode.Sets()
}
