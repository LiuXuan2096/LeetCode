package Go_LeetCode

// 这道题没做完，没看懂讲解 先不做了。
import "sort"

type Node_218 struct {
	x     int
	isAdd bool
	h     int
}

type NodeSlice []Node_218

func (ns NodeSlice) Len() int {
	return len(ns)
}

func (ns NodeSlice) Less(i, j int) bool {
	return ns[i].x < ns[j].x
}

func (ns NodeSlice) Swap(i, j int) {
	ns[i], ns[j] = ns[j], ns[i]
}

func getSkyline(buildings [][]int) [][]int {
	nodes := make([]Node_218, len(buildings)*2)
	for i := 0; i < len(buildings); i++ {
		nodes[i*2] = Node_218{buildings[i][0], true, buildings[i][2]}
		nodes[i*2+1] = Node_218{buildings[i][1], false, buildings[i][2]}

	}
	sort.Sort(NodeSlice(nodes))

	mapHeightTimes := make(map[int]int)
	//mapXHeight := make(map[int]int)

	for i := 0; i < len(nodes); i++ {
		if nodes[i].isAdd {
			mapHeightTimes[nodes[i].h]++
		} else {
			if mapHeightTimes[nodes[i].h] == 1 {
				//delete(mapHeightTimes)
			}
		}
	}

	return [][]int{}
}

func max_218(a, b int) int {
	if a > b {
		return a
	}
	return b
}
