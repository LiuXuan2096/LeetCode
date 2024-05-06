package Go_LeetCode

/*
https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
*/

func findTargetIn2DPlants(plants [][]int, target int) bool {
	if plants == nil || len(plants) == 0 || len(plants[0]) == 0 {
		return false
	}
	row := 0
	col := len(plants[0]) - 1

	for row < len(plants) && col >= 0 {
		if plants[row][col] == target {
			return true
		} else if plants[row][col] > target {
			col--
		} else {
			row++
		}
	}
	return false
}
